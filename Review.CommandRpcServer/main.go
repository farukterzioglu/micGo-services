package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	pb "github.com/farukterzioglu/micGo-services/Review.CommandRpcServer/reviewservice"
	"google.golang.org/grpc"
)

func main() {
	flag.Parse()

	lis, err := net.Listen("tcp", ":3000")
	if err != nil {
		log.Fatalf("failed to listen: %v\n", err)
	}
	fmt.Printf("Running server at %s...\n", ":3000")

	grpcServer := grpc.NewServer()
	pb.RegisterReviewServiceServer(grpcServer, NewCommandServer())
	grpcServer.Serve(lis)
}
