# Start Kafka & Zookeper

https://github.com/ches/docker-kafka

`Set-Variable -Name "ipAddress" -Value "192.168.1.5"`

`docker run -d --name zookeeper -p 2181:2181 jplock/zookeeper`

`docker run -d --name kafka -p 7203:7203 -p 9092:9092 -e KAFKA_ADVERTISED_HOST_NAME=$ipAddress -e KAFKA_MESSAGE_MAX_BYTES=3000000 -e KAFKA_REPLICA_FETCH_MAX_BYTES=3100000 -e ZOOKEEPER_IP=$ipAddress ches/kafka`

`docker run --rm ches/kafka kafka-topics.sh --create --topic tags --replication-factor 1 --partitions 1 --zookeeper ($ipAddress + ':2181')`

`docker run --rm ches/kafka kafka-topics.sh --list --zookeeper ($ipAddress + ':2181')`

`docker run --rm --interactive ches/kafka kafka-console-producer.sh --topic tags --broker-list ($ipAddress + ':9092')`

(From another powershell instance)  
`Set-Variable -Name "ipAddress" -Value "192.168.1.5"`

`docker run --rm ches/kafka kafka-console-consumer.sh --topic create-review --from-beginning --zookeeper ($ipAddress + ':2181')`

`docker run --rm --interactive ches/kafka kafka-consumer-groups.sh --new-consumer --describe --group group1 --bootstrap-server (\$ipAddress + ':9092')`

# Check kafka instance details

`docker exec -it zookeeper bash`  
`bin/zkCli.sh -server 127.0.0.1:2181`  
`ls /brokers`  
`ls /brokers/topics`  
`ls /consumers`

# Consumer with Golang

`cd .\KafkaComparer.Consumer.Golang`  
`docker build -t kafkacomparerconsumer:go .`  
`docker run -it kafkacomparerconsumer:go`

# Producer with Golang

`cd .\KafkaComparer.Producer.Golang`  
`docker build -t kafkacomparerproducer:go .`  
`docker run -it kafkacomparerproducer:go`

Sample message to send  
`{ "review": { "text": "Liked it!!!", "star": 5 } }`

# Compare Engine

Reads from kafka topic and handles commands (new comment etc.) in go routines  
`docker build -f .\build\Review.CommandEngine\Dockerfile -t command-engine:latest .`  
`docker run -it command-engine:latest`

# Command Rpc Server

Handles rpc commands
`docker build -f .\build\Review.CommandRpcServer\Dockerfile -t command-rpcserver:latest .`  
`docker run -it command-command-rpcserver:latest`
