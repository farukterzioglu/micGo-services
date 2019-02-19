ARG GO_VERSION=1.11
FROM golang:${GO_VERSION} AS builder

WORKDIR /src
COPY ./Review.API/go.mod ./Review.API/go.sum ./
RUN go mod download

COPY ./Review.API ./

ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on GOFLAGS=-mod=vendor 
RUN go build -a -o app .

RUN set -x && go get github.com/go-swagger/go-swagger/cmd/swagger
RUN swagger generate spec -o ./swaggerui/swagger.json --scan-models

FROM alpine:latest
RUN apk add --no-cache bash
RUN mkdir /app
WORKDIR /app
COPY --from=builder /src/Review.API/app .
COPY --from=builder /go/Review.API/swaggerui/ ./swaggerui

ENTRYPOINT ["./app"]