# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOVET=$(GOCMD) vet
GOFMT=gofmt
GOLINT=golint
BINARY_NAME=fizzbuzz

all: lint tool test build

test: 
	export GIN_MODE=release && $(GOTEST) -short $(go list ./... | grep -v /vendor/)

build:
	$(GOBUILD) -v .

tool:
	$(GOVET) ./...; true
	$(GOFMT) -w .

coverage:
	scripts/coverage.sh

clean:
	go clean -i .
	rm -f $(BINARY_NAME)

lint:
	$(GOLINT) -set_exit_status $($(GOCMD) list ./... | grep -v /vendor/)

deps:
	$(GOGET) github.com/golang/dep/cmd/dep
	$(GOGET) golang.org/x/lint/golint
	dep ensure

docker-build: clean deps build
	docker build -t fizzbuzz .
	docker run --rm -e GIN_MODE='release' -p 8080:8080 fizzbuzz:latest 

help:
	@echo "make: compile packages and dependencies"
	@echo "make tool: run specified go tool"
	@echo "make lint: golint ./..."
	@echo "make clean: remove object files and cached files"
	@echo "make deps: get the deployment tools"
	@echo "make coverage: get the coverage of my files"
	@echo "make docker-build: build a docker image and run the container"