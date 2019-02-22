ARG GO_VERSION=1.11
FROM golang:${GO_VERSION} AS builder

ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on 

WORKDIR /src

COPY ./go.mod ./go.sum ./
RUN go mod download

COPY ./Review.Domain ./Review.Domain
COPY ./Review.CommandRpcServer ./Review.CommandRpcServer

WORKDIR /src/Review.CommandRpcServer
RUN go build -a -o app .


FROM scratch
WORKDIR /root/
COPY --from=builder /src/Review.CommandRpcServer/app .

ENTRYPOINT ["./app"]
EXPOSE 3000