.PHONY: all build run

PROGRAM=server.out
CONTAINER=api.moki.codes:latest

all: build run

build:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o $(PROGRAM) .
	docker build -t $(CONTAINER) -f Dockerfile .
run:
	docker run -t $(CONTAINER)
