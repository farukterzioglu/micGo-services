FROM golang:1.10.1 AS builder
RUN go version

RUN set -x && \
    go get github.com/bsm/sarama-cluster && \
    go get google.golang.org/grpc && \
    go get github.com/golang/protobuf/proto

COPY ./Review.Domain /go/src/github.com/farukterzioglu/micGo-services/Review.Domain/
COPY ./Review.CommandRpcServer /go/src/github.com/farukterzioglu/micGo-services/Review.CommandRpcServer/
COPY ./Review.CommandEngine /go/src/github.com/farukterzioglu/micGo-services/Review.CommandEngine/
WORKDIR /go/src/github.com/farukterzioglu/micGo-services/Review.CommandEngine/ 

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o app .

FROM alpine:latest
RUN apk add --no-cache bash
RUN mkdir /app
WORKDIR /app
COPY --from=builder /go/src/github.com/farukterzioglu/micGo-services/Review.CommandEngine/app .

ENTRYPOINT ["./app"]