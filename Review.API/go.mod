module Review.API

require (
	github.com/DataDog/zstd v1.3.5 // indirect
	github.com/Shopify/sarama v1.20.1
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/eapache/go-resiliency v1.1.0 // indirect
	github.com/eapache/go-xerial-snappy v0.0.0-20180814174437-776d5712da21 // indirect
	github.com/eapache/queue v1.1.0 // indirect
	github.com/farukterzioglu/micGo-services v0.0.0-20190213202828-9b2573277f69
	github.com/go-kit/kit v0.8.0
	github.com/go-logfmt/logfmt v0.4.0 // indirect
	github.com/golang/protobuf v1.2.1-0.20190205222052-c823c79ea157 // indirect
	github.com/golang/snappy v0.0.0-20190218232222-2a8bb927dd31 // indirect
	github.com/google/uuid v1.1.0 // indirect
	github.com/gorilla/mux v1.7.0
	github.com/pierrec/lz4 v2.0.5+incompatible // indirect
	github.com/rcrowley/go-metrics v0.0.0-20181016184325-3113b8401b8a // indirect
	google.golang.org/grpc v1.18.0
)

replace github.com/farukterzioglu/micGo-services/Review.CommandRpcServer v0.0.0-20190219084034-b730cfdf7f8b => ../Review.CommandRpcServer
