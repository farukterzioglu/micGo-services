package commandhandlers

import "context"
import "github.com/farukterzioglu/micGo-services/Review.Domain/Commands/V1"

// HandlerRequest request model for handlers
type HandlerRequest struct {
	Command         commands.ICommand
	HandlerResponse chan interface{}
	ErrResponse     chan error
}

// ICommandHandler interface for command handlers
type ICommandHandler interface {
	HandleAsync(ctx context.Context, request HandlerRequest)
}
