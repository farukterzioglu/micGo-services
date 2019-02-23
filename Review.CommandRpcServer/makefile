.DEFAULT_GOAL := help

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=command-rpcserver
BINARY_UNIX=$(BINARY_NAME)_unix

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}' 
## help output can be sorted alphabetically by addin :  | sort | awk 

all: test build
build: ## Builds the project
	$(GOBUILD) -o $(BINARY_NAME) -v
build-linux: ## Builds the project for linux
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_UNIX) -v
test: ## Tests go packages 
	$(GOTEST) -v ./...
clean: ## Cleans build artifacts
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_UNIX)
run: ## Runs the app
	$(GOBUILD) -o $(BINARY_NAME) -v
	./$(BINARY_NAME)
deps: ## Download dependencies
	dep ensure
d-build: ## Builds Dockerfile
	docker build -f ../build/Review.CommandRpcServer/Dockerfile -t command-rpcserver:latest ../.
d-run: ## Runs built docker image 
	docker run -it -p 3000:3000 command-rpcserver:latest