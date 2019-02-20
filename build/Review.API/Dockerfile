ARG GO_VERSION=1.11
FROM golang:${GO_VERSION} AS builder

WORKDIR /src
COPY ./Review.API/go.mod ./Review.API/go.sum ./
RUN go mod download
RUN go get github.com/golang/protobuf@master

COPY ./Review.API ./

ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on 
RUN go build -a -o app .

# TODO : Activate after those issues fixed
# https://github.com/go-swagger/go-swagger/issues/1681
# https://github.com/go-swagger/go-swagger/issues/1671
# RUN set -x && go get github.com/go-swagger/go-swagger/cmd/swagger
# RUN swagger generate spec -o ./swaggerui/swagger.json --scan-models

FROM alpine:latest
RUN apk add --no-cache bash
RUN mkdir /app
WORKDIR /app
COPY --from=builder /src/app .
COPY --from=builder /src/swaggerui/ ./swaggerui

ENTRYPOINT ["./app"]