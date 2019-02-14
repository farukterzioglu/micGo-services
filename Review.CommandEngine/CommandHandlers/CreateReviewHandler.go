package commandhandlers

import (
	"context"
	"fmt"
	"log"

	"github.com/afex/hystrix-go/hystrix"
	pb "github.com/farukterzioglu/micGo-services/Review.CommandRpcServer/reviewservice"
	"github.com/farukterzioglu/micGo-services/Review.Domain/Commands/V1"
	"github.com/farukterzioglu/micGo-services/Review.Domain/Models"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

// CreateReviewHandler is the handler for CreateReview command
type CreateReviewHandler struct {
	client pb.ReviewServiceClient
}

// NewCreateReviewHandler creates and returns new 'create review' command handler
func NewCreateReviewHandler(c *pb.ReviewServiceClient) *CreateReviewHandler {
	hystrix.ConfigureCommand("save-review-rpc", hystrix.CommandConfig{Timeout: 5000})

	return &CreateReviewHandler{
		client: *c,
	}
}

func populateRPCCommand(review *models.Review) *pb.NewReviewRequest {
	return &pb.NewReviewRequest{
		Review: &pb.Review{
			ReviewID:  review.ID,
			ProductID: review.ProductID,
			UserID:    review.UserID,
			Text:      review.Text,
			Star:      int32(review.Star),
		},
	}
}

// HandleAsync handles string message
func (handler *CreateReviewHandler) HandleAsync(ctx context.Context, request HandlerRequest) {
	var createReviewCommand commands.CreateReviewCommand
	createReviewCommand = request.Command.(commands.CreateReviewCommand)

	ctx = metadata.NewOutgoingContext(
		ctx,
		metadata.Pairs("request-owner", "CreateReviewHandler"),
	)
	// metadata.AppendToOutgoingContext(ctx, "key", "value")

	hystrix.Go("save-review-rpc", func() error {
		response, err := handler.client.SaveReview(ctx, populateRPCCommand(&createReviewCommand.Review))
		if err != nil {
			return err
		}

		fmt.Printf("Created a review with id : %s \n", response.ReviewId)
		ctx = models.NewContextWithReviewID(ctx, response.ReviewId)

		request.HandlerResponse <- response.ReviewId

		return nil
	}, func(err error) error {
		errStatus, _ := status.FromError(err)
		log.Printf("rpc error : %v\n", errStatus)

		request.ErrResponse <- err
		return nil
	})
}
