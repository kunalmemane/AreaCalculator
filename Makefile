.PHONY: build test clean run help all docker-build docker-push oc-deploy oc-delete check-podman-login check-oc-login changelog

APP_NAME = area-calculator
IMAGE_REGISTRY = quay.io
IMAGE_TAG = $(IMAGE_REGISTRY)/kmemane/$(APP_NAME)

help:
	@echo "Make - Area Calculator:"
	@echo "Available targets:"
	@echo "  build    		- Build the application"
	@echo "  test     		- Run tests"
	@echo "  clean    		- Clean the project"
	@echo "  run      		- Build and run the application"
	@echo "  all	  		- Clean, Test, build and run the application"
	@echo "  podman-build   	- Build Application Image using Podman" 
	@echo "  podman-run		- Run Application Container Image"
	@echo "  podman-push    	- Push Application Image to quay.io [login required]"
	@echo "  oc-deploy    		- Deploy application to openshift using docker strategy [login required]"
	@echo "  help     		- Display this help message"

build:
	CGO_ENABLED=0  go build -o ./bin/main cmd/main.go

test:
	go test -v ./...

clean:
	go clean
	rm -f ./bin/main

run: build
	./bin/main

all:
	make clean
	make test
	make build
	make run

changelog:
	@./script/changelog.sh

# Podman recipes
check-podman-login:
	@if podman login $(IMAGE_REGISTRY) --get-login; \
	then \
		echo "You are logged in to $(IMAGE_REGISTRY)"; \
	else \
		echo "You are not logged in to $(IMAGE_REGISTRY)"; \
		exit 1;\
	fi
		
podman-build: 
	podman build -f Dockerfile -t $(IMAGE_TAG)

podman-run: podman-build
	podman run -p 8080:8080 $(IMAGE_TAG)

podman-push: check-podman-login podman-build
	podman push $(IMAGE_TAG)


# Openshift (oc) recipes
check-oc-login:
	@if oc whoami; \
	then \
		echo "You are logged in to Openshift cluster as $$(oc whoami)"; \
	else \
		echo "You are not logged in to Openshift cluster"; \
		exit 1; \
	fi

oc-deploy: check-oc-login
	oc project $(APP_NAME)
	oc new-app . --name $(APP_NAME) --strategy=docker
	@sleep 5
	oc start-build $(APP_NAME) --from-dir=./ --follow=true --wait=true
	oc expose svc $(APP_NAME)
	oc get routes.route.openshift.io

oc-delete:
	oc project $(APP_NAME)
	oc delete all --selector=app=${APP_NAME}