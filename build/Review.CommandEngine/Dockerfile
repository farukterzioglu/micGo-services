ARG GO_VERSION=1.11
FROM golang:${GO_VERSION} AS builder

ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on 

WORKDIR /src

COPY ./go.mod ./go.sum ./
RUN go mod download
# go get github.com/golang/protobuf/proto

COPY ./Review.Domain ./Review.Domain/
COPY ./Review.CommandRpcServer ./Review.CommandRpcServer/
COPY ./Review.CommandEngine ./Review.CommandEngine/

WORKDIR /src/Review.CommandEngine/ 
RUN go build -a -o app .

FROM alpine:latest
RUN apk add --no-cache bash
RUN mkdir /app
WORKDIR /app
COPY --from=builder /src/Review.CommandEngine/app .

ENTRYPOINT ["./app"]