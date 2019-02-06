Generate proto service file

```
cd ./Review.CommandRpcServer

protoc -I reviewservice/ reviewservice/review_service.proto --go_out=plugins=grpc:reviewservice
```

https://blog.oklahome.net/2018/07/protoactor-go-introduction.html
