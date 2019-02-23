.PHONY : build-webapi build-commanengine build-rpcserver run-webapi run-commandengine run-rpcserver
.DEFAULT_GOAL := help

DOCKERBUILD=docker build 
DOCKERRUN=docker run
HOSTIP="172.19.0.2:9092"

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}' 
## help output can be sorted alphabetically by addin :  | sort | awk 

build-all: build-webapi build-commanengine build-rpcserver ## Builds all docker files
run-all: run-rpcserver run-webapi run-commandengine

kafka-up: ## Starts Kafka on docker compose (along with Zookeeper)
	@echo "\033[36mStarting Kafka...\033[0m"
	@docker-compose up

build-webapi: ## Builds Review Wep API dockerfile
	@echo "\033[36mBuilding web api...\033[0m"
	$(DOCKERBUILD) -f ./build/Review.API/Dockerfile -t review-api:latest .
run-webapi: ## Start Review Web API on docker
	$(DOCKERRUN) -it -p 8000:8000 review-api:latest -kafka_brokers=$(HOSTIP) -server_addr="localhost:3000" -port="8000"

build-commanengine: ## Builds Command Engine dockerfile
	@echo "\033[36mBuilding command engine...\033[0m"
	$(DOCKERBUILD) -f ./build/Review.CommandEngine/Dockerfile -t command-engine:latest .
run-commandengine: ## Starts Command Engine
	$(DOCKERRUN) -it command-engine:latest -server_addr="localhost:3000" -kafka_brokers="$(HOSTIP):9092" -group_id="test"

build-rpcserver: ## Builds RPC Server dockerfile
	@echo "\033[36mBuilding rpc server...\033[0m"
	$(DOCKERBUILD) -f ./build/Review.CommandRpcServer/Dockerfile -t command-rpcserver:latest .
run-rpcserver: ## Starts RPC Server
	$(DOCKERBUILD) -it -p 3000:3000 command-rpcserver:latest

