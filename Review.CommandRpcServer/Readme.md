Generate proto service file

```
cd ./Review.CommandRpcServer

// TODO : Install version2 protoc or change 'proto.ProtoPackageIsVersion3' to 'proto.ProtoPackageIsVersion2'
// dep (dependency manager) downloads version 2 proto buffers but protoc generates version 3 codes
// so couldn't build review api project. Needed to make manuel changes in generated code.
protoc -I reviewservice/ reviewservice/review_service.proto --go_out=plugins=grpc:reviewservice
```

https://blog.oklahome.net/2018/07/protoactor-go-introduction.html
