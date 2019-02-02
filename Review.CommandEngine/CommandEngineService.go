package main

import (
	"context"
	"encoding/json"

	"github.com/farukterzioglu/micGo-services/Review.CommandEngine/CommandHandlers"
	pb "github.com/farukterzioglu/micGo-services/Review.CommandRpcServer/reviewservice"
	"github.com/farukterzioglu/micGo-services/Review.Domain/Commands/V1"
	"github.com/farukterzioglu/micGo-services/Review.Domain/Models"
)

// CommandRequest is the request type for commands
type CommandRequest struct {
	CommandMessage models.CommandMessage
	ResponseCh     chan interface{}
	ErrCh          chan error
}

type commandCreatorFunc func() commandhandlers.ICommandHandler

// CommandEngineService is service that handles command messages
type CommandEngineService struct {
	client *pb.ReviewServiceClient
}

// NewCommandEngineService returns new command engine service
func NewCommandEngineService(c *pb.ReviewServiceClient) *CommandEngineService {
	return &CommandEngineService{
		client: c,
	}
}

// HandleMessage handles consumed command message
func (service *CommandEngineService) HandleMessage(ctx context.Context, request CommandRequest) {
	commandMessage := request.CommandMessage
	// fmt.Fprintf(os.Stdout, "%s/%d/%d\t%s\t%s\n", msg.Topic, msg.Partition, msg.Offset, msg.Key, msg.Value)

	// Handler
	var handler commandhandlers.ICommandHandler
	var command commands.ICommand

	switch commandMessage.CommandType {
	case "create-review":
		var createReviewCommand commands.CreateReviewCommand
		json.Unmarshal(commandMessage.CommandData, &createReviewCommand)

		handler = commandhandlers.NewCreateReviewHandler(service.client)
		command = createReviewCommand
	case "rate-review":
		var rateReviewCommand commands.RateReviewCommand
		json.Unmarshal(commandMessage.CommandData, &rateReviewCommand)

		handler = commandhandlers.NewRateReviewHandler(service.client)
		command = rateReviewCommand
	default:
		handler = commandhandlers.NewDefaultHandler()
		command = commandMessage.CommandData
	}

	// Request
	var handlerRequest commandhandlers.HandlerRequest
	handlerRequest = commandhandlers.HandlerRequest{
		Command:         command,
		HandlerResponse: request.ResponseCh,
		ErrResponse:     request.ErrCh,
	}

	handler.HandleAsync(ctx, handlerRequest)
}
