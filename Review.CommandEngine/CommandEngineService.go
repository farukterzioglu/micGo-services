package main

import (
	"context"

	"github.com/Shopify/sarama"
	"github.com/farukterzioglu/micGo-services/Review.CommandEngine/CommandHandlers"

	pb "github.com/farukterzioglu/micGo-services/Review.CommandRpcServer/reviewservice"
)

// CommandRequest is the request type for commands
type CommandRequest struct {
	Msg        *sarama.ConsumerMessage
	ResponseCh chan interface{}
	ErrCh      chan error
}

type commandCreatorFunc func() commandhandlers.ICommandHandler

var commandMap map[string]commandCreatorFunc

// CommandEngineService is service that handles command messages
type CommandEngineService struct {
	client *pb.ReviewServiceClient
}

// NewCommandEngineService returns new command engine service
func NewCommandEngineService(c *pb.ReviewServiceClient) *CommandEngineService {
	commandMap = make(map[string]commandCreatorFunc)
	commandMap["create-review"] = func() commandhandlers.ICommandHandler {
		return commandhandlers.NewCreateReviewHandler(c)
	}
	commandMap["rate-review"] = func() commandhandlers.ICommandHandler {
		return commandhandlers.NewRateReviewHandler()
	}

	return &CommandEngineService{
		client: c,
	}
}

func (service *CommandEngineService) getTopicList() []string {
	keys := make([]string, len(commandMap))

	i := 0
	for k := range commandMap {
		keys[i] = k
		i++
	}

	return keys
}

// HandleMessage handles consumed command message
func (service *CommandEngineService) HandleMessage(ctx context.Context, request CommandRequest) {
	msg := request.Msg
	// fmt.Fprintf(os.Stdout, "%s/%d/%d\t%s\t%s\n", msg.Topic, msg.Partition, msg.Offset, msg.Key, msg.Value)

	// Request
	var handlerRequest commandhandlers.HandlerRequest
	handlerRequest = commandhandlers.HandlerRequest{
		Command:         msg.Value,
		HandlerResponse: request.ResponseCh,
		ErrResponse:     request.ErrCh,
	}

	// Handler
	var handler commandhandlers.ICommandHandler
	if createHandler, ok := commandMap[msg.Topic]; ok {
		handler = createHandler()
	} else {
		handler = commandhandlers.NewDefaultHandler()
	}
	handler.HandleAsync(ctx, handlerRequest)
}
