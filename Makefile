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
	podman build -f Dockerfile -t quay.io/rh-ee-kmemane/area-calculator

podman-run:
	podman run -p 8080:8080 quay.io/rh-ee-kmemane/area-calculator

podman-push:
	podman push quay.io/rh-ee-kmemane/area-calculator

oc-deploy:
	oc new-app . --name area-calculator --strategy=docker

help:
	@echo "Make - area calculator:"
	@echo "Available targets:"
	@echo "  build    - Build the application"
	@echo "  test     - Run tests"
	@echo "  clean    - Clean the project"
	@echo "  run      - Build and run the application"
	@echo "  help     - Display this help message"

.PHONY: build test clean run help all docker-build docker-push