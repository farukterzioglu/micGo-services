FROM golang:1.10.1 AS builder
RUN go version

RUN set -x && \
    go get google.golang.org/grpc && \
    go get github.com/AsynkronIT/protoactor-go/actor

COPY ./Review.Domain /go/src/github.com/farukterzioglu/micGo-services/Review.Domain/
COPY ./Review.CommandRpcServer /go/src/github.com/farukterzioglu/micGo-services/Review.CommandRpcServer/
WORKDIR /go/src/github.com/farukterzioglu/micGo-services/Review.CommandRpcServer/

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o app .

FROM scratch
WORKDIR /root/
COPY --from=builder /go/src/github.com/farukterzioglu/micGo-services/Review.CommandRpcServer/app .

ENTRYPOINT ["./app"]
EXPOSE 3000