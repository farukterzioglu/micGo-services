FROM golang:1.10.1 AS builder
RUN go version

COPY ./Review.CommandEngine /go/src/github.com/farukterzioglu/micGo-services/Review.CommandRpcServer/
WORKDIR /go/src/github.com/farukterzioglu/micGo-services/Review.CommandRpcServer/
RUN set -x && \
    go get github.com/bsm/sarama-cluster 

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o app .

FROM scratch
WORKDIR /root/
COPY --from=builder /go/src/github.com/farukterzioglu/micGo-services/Review.CommandRpcServer/app .

ENTRYPOINT ["./app"]