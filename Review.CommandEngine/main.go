package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/Shopify/sarama"
	cluster "github.com/bsm/sarama-cluster"
	pb "github.com/farukterzioglu/micGo-services/Review.CommandRpcServer/reviewservice"
	"github.com/farukterzioglu/micGo-services/Review.Domain/Models"
	"google.golang.org/grpc"
)

var (
	serverAddr   = flag.String("server_addr", "127.0.0.1:10000", "The rpc server address in the format of host:port")
	kafkaBrokers = flag.String("kafka_brokers", "127.0.0.1:9092", "The kafka broker address in the format of host:port")
	groupID      = flag.String("group_id", "review-command-engine", "Group ID")
)

func main() {
	flag.Parse()
	fmt.Printf("Broker address : %s\n", *kafkaBrokers)
	fmt.Printf("RPC server address : %s\n", *serverAddr)
	fmt.Printf("Group Id : %s\n", *groupID)

	// Configure gRpc client
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())

	conn, err := grpc.Dial(*serverAddr, opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
		panic(err)
	}
	defer conn.Close()

	rpcConnState := conn.GetState()
	fmt.Printf("Rpc cpnnection state : %s\n", rpcConnState)

	client := pb.NewReviewServiceClient(conn)

	// Configure command engine service
	var commandEngineService *CommandEngineService
	commandEngineService = NewCommandEngineService(&client)

	// init (custom) config, enable errors and notifications
	config := cluster.NewConfig()
	config.Consumer.Return.Errors = true
	config.Group.Return.Notifications = true
	// Config.Consumer.Offsets.Initial = sarama.OffsetOldest

	// init kafka consumer
	brokers := []string{*kafkaBrokers} // TODO : parse comma seperated broker list
	topics := []string{"review-commands"}
	consumer, err := cluster.NewConsumer(brokers, *groupID, topics, config)
	if err != nil {
		panic(err)
	}
	defer consumer.Close()

	// trap SIGINT to trigger a shutdown.
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	// consume errors
	go func() {
		for err := range consumer.Errors() {
			log.Printf("Error: %s\n", err.Error())
		}
	}()

	// consume notifications
	go func() {
		for ntf := range consumer.Notifications() {
			log.Printf("Rebalanced: %+v\n", ntf)
		}
	}()

	// consume failed messages
	failedMsgChn := make(chan *sarama.ConsumerMessage)
	go func(failedMsgCh chan *sarama.ConsumerMessage) {
		for msg := range failedMsgChn {
			fmt.Printf("Message failed : %s\n", msg.Value)
			// TODO : Process failed messages again
		}
	}(failedMsgChn)

	// handler function for consumer messages
	var wg sync.WaitGroup
	handlerFunc := func(msg *sarama.ConsumerMessage) {
		wg.Add(1)
		go func() {
			defer wg.Done()
			defer func() {
				if r := recover(); r != nil {
					fmt.Printf("Panic: %+v\n", r)
					failedMsgChn <- msg
				}
			}()

			var (
				ctx    context.Context
				cancel context.CancelFunc
			)
			ctx, cancel = context.WithCancel(context.Background())
			defer cancel()

			var v models.CommandMessage
			json.Unmarshal(msg.Value, &v)

			request := CommandRequest{
				CommandMessage: v,
				ResponseCh:     make(chan interface{}),
				ErrCh:          make(chan error),
			}

			go commandEngineService.HandleMessage(ctx, request)

		Completed:
			for {
				select {
				case resp := <-request.ResponseCh:
					returnValue := resp.(string)
					fmt.Printf("Review id : %s\n", returnValue)
					// reviewID := models.ReviewIDFromContext(ctx)
					// fmt.Printf("Review id from context: %s\n", reviewID)
					break Completed
				case err := <-request.ErrCh:
					fmt.Printf("Request failed : %s\n", err.Error())
					// TODO : Retry
					failedMsgChn <- msg
					break Completed
				case <-ctx.Done():
					fmt.Printf("Request failed : %s\n", ctx.Err())
					failedMsgChn <- msg
					break Completed
				case <-time.After(time.Minute):
					fmt.Printf("Request timedout!\n")
					failedMsgChn <- msg
					cancel()
					break Completed
				}
			}

			consumer.MarkOffset(msg, "")
		}()
	}

	// channel for handling messsaged
	msgch := make(chan *sarama.ConsumerMessage)

	go func(channel chan *sarama.ConsumerMessage) {
		for newMsg := range channel {
			handlerFunc(newMsg)
		}
		wg.Wait()
	}(msgch)

	// consume messages, watch signals
	for {
		select {
		case msg, ok := <-consumer.Messages():
			if ok {
				msgch <- msg
			}
		case <-signals:
			close(msgch)
			return
		}
	}
}
