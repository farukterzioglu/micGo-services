package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/AsynkronIT/protoactor-go/actor"
	"github.com/farukterzioglu/micGo-services/Review.CommandRpcServer/actors"
	pb "github.com/farukterzioglu/micGo-services/Review.CommandRpcServer/reviewservice"
	"google.golang.org/grpc"
)

func main() {
	flag.Parse()

	lis, err := net.Listen("tcp", "localhost:10000")
	if err != nil {
		log.Fatalf("failed to listen: %v\n", err)
	}
	fmt.Printf("Running server at %s...\n", "localhost:10000")

	reviewPid := StartReviewActor()

	grpcServer := grpc.NewServer()
	pb.RegisterReviewServiceServer(grpcServer, NewCommandServer(reviewPid))
	grpcServer.Serve(lis)
}

// StartReviewActor ...
func StartReviewActor() *actor.PID {
	decider := func(reason interface{}) actor.Directive {
		fmt.Println("handling failure for child")
		return actor.StopDirective
	}
	supervisor := actor.NewOneForOneStrategy(10, 1000, decider)
	props := actor.
		FromProducer(actors.NewReviewActor).
		WithSupervisor(supervisor)

	reviewPid := actor.Spawn(props)
	return reviewPid

	// ordersProps := actor.FromProducer(func() actor.Actor {
	// 	return &actors.OrdersActor{}
	// })
	// ordersPid := actor.Spawn(ordersProps)

	// reviewProps := actor.FromProducer(func() actor.Actor {
	// 	return &actors.ReviewActor{
	// 		ordersPid : ordersPid
	// 	}
	// })
	// reviewPid := actor.Spawn(reviewProps)

}
