all: build docker-image docker-run
.PHONY: all

build:
	CGO_ENABLED=0 GO111MODULES=on go build .
.PHONY: build

docker-image:
	docker build -t http-echo-request-server:latest .
.PHONY: image

docker-run:
	docker run -p "8080:8080" http-echo-request-server:latest
.PHONY: run-docker
