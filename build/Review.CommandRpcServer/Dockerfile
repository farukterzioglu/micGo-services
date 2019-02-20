ARG GO_VERSION=1.11
FROM golang:${GO_VERSION} AS builder

WORKDIR /src
COPY ./Review.CommandRpcServer/go.mod ./Review.CommandRpcServer/go.sum ./
RUN go mod download

COPY ./Review.API ./

ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on 
RUN go build -a -o app .

# go get google.golang.org/grpc && \
# go get github.com/AsynkronIT/protoactor-go/actor

COPY ./Review.Domain /go/src/github.com/farukterzioglu/micGo-services/Review.Domain/
COPY ./Review.CommandRpcServer /go/src/github.com/farukterzioglu/micGo-services/Review.CommandRpcServer/
WORKDIR /go/src/github.com/farukterzioglu/micGo-services/Review.CommandRpcServer/

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o app .

FROM scratch
WORKDIR /root/
COPY --from=builder /go/src/github.com/farukterzioglu/micGo-services/Review.CommandRpcServer/app .

ENTRYPOINT ["./app"]
EXPOSE 3000