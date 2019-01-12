FROM golang:1.10.1 AS builder
RUN go version

RUN set -x && \
    go get github.com/Shopify/sarama 

COPY ./KafkaComparer.Producer.Golang /go/src/github.com/farukterzioglu/micGo-services/KafkaComparer.Producer.Golang/
WORKDIR /go/src/github.com/farukterzioglu/micGo-services/KafkaComparer.Producer.Golang/

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o app .

FROM scratch
WORKDIR /root/
COPY --from=builder /go/src/github.com/farukterzioglu/micGo-services/KafkaComparer.Producer.Golang/app .

ENTRYPOINT ["./app"]