.PHONY: all build build-container run-container

PROGRAM=server.out
CONTAINER=api.moki.codes:latest
DSTPORT=4000
SRCPORT=80

all: clean build build-container run-container

build:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o $(PROGRAM) .
build-container:
	docker build -t $(CONTAINER) -f Dockerfile .
clean:
	rm ./*.out
run-container:
	docker run --publish $(DSTPORT):$(SRCPORT) -t $(CONTAINER)
