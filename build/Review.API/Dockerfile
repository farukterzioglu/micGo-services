FROM golang:1.10.1 AS builder

RUN set -x && \
    go get github.com/Shopify/sarama && \
    go get github.com/gorilla/mux && \
    go get github.com/go-swagger/go-swagger/cmd/swagger && \
    go get github.com/golang/protobuf/proto && \
    go get google.golang.org/grpc

# TODO : Get all deps
RUN set -x && \
    go get github.com/google/uuid 

COPY ./Review.Domain /go/src/github.com/farukterzioglu/micGo-services/Review.Domain/
COPY ./Review.API /go/src/github.com/farukterzioglu/micGo-services/Review.API/
COPY ./Review.CommandRpcServer /go/src/github.com/farukterzioglu/micGo-services/Review.CommandRpcServer/

WORKDIR /go/src/github.com/farukterzioglu/micGo-services/Review.API/ 

RUN swagger generate spec -o ./swaggerui/swagger.json --scan-models
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o app .

FROM alpine:latest
RUN apk add --no-cache bash
RUN mkdir /app
WORKDIR /app
COPY --from=builder /go/src/github.com/farukterzioglu/micGo-services/Review.API/app .
COPY --from=builder /go/src/github.com/farukterzioglu/micGo-services/Review.API/swaggerui/ ./swaggerui

ENTRYPOINT ["./app"]