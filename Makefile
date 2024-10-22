APP_NAME = area-calculator
IMAGE_TAG = quay.io/rh-ee-kmemane/${APP_NAME}

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

podman-build:
	podman build -f Dockerfile -t ${IMAGE_TAG}

podman-run:
	podman run -p 8080:8080 ${IMAGE_TAG}

podman-push:
	podman push ${IMAGE_TAG}

oc-deploy:
	oc project ${APP_NAME}
	oc new-app . --name ${APP_NAME} --strategy=docker
	sleep 5
	oc start-build ${APP_NAME} --from-dir=./ --follow=true --wait=true
	oc expose svc ${APP_NAME}
	oc get routes.route.openshift.io

oc-delete:
	oc project ${APP_NAME}
	oc delete all --all

.PHONY: build test clean run help all docker-build docker-push