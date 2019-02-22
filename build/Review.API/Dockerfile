ARG GO_VERSION=1.11
FROM golang:${GO_VERSION} AS builder

ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64

# swagger doesn't support go modules yet.
# https://github.com/go-swagger/go-swagger/issues/1681
# https://github.com/go-swagger/go-swagger/issues/1671
ENV GO111MODULE=off 

COPY ./Review.API/Gopkg.toml ./Review.API/Gopkg.lock /go/src/github.com/farukterzioglu/micGo-services/Review.API/
WORKDIR /go/src/github.com/farukterzioglu/micGo-services/Review.API/ 

RUN go get -u github.com/golang/dep/...
RUN dep ensure --vendor-only

COPY ./Review.Domain /go/src/github.com/farukterzioglu/micGo-services/Review.Domain/
COPY ./Review.API /go/src/github.com/farukterzioglu/micGo-services/Review.API/
COPY ./Review.CommandRpcServer /go/src/github.com/farukterzioglu/micGo-services/Review.CommandRpcServer/

RUN go build -a -o app .
# RUN swagger generate spec -o ./swaggerui/swagger.json --scan-models

FROM alpine:latest
RUN apk add --no-cache bash
RUN mkdir /app
WORKDIR /app
COPY --from=builder /go/src/github.com/farukterzioglu/micGo-services/Review.API/app .
# COPY --from=builder /go/src/github.com/farukterzioglu/micGo-services/Review.API/swaggerui/ ./swaggerui

ENTRYPOINT ["./app"]