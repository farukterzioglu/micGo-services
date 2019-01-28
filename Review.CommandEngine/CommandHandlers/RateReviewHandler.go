package commandhandlers

import (
	"context"
	"encoding/json"
	"fmt"

	pb "github.com/farukterzioglu/micGo-services/Review.CommandRpcServer/reviewservice"
	"github.com/farukterzioglu/micGo-services/Review.Domain/Commands/V1"
)

// RateReviewHandler is the handler for CreateReview command
type RateReviewHandler struct {
	client pb.ReviewServiceClient
}

// NewRateReviewHandler creates and returns new 'rate review' command handler
func NewRateReviewHandler(c *pb.ReviewServiceClient) *RateReviewHandler {
	return &RateReviewHandler{
		client: *c,
	}
}

// HandleAsync handles string message
func (handler *RateReviewHandler) HandleAsync(ctx context.Context, request HandlerRequest) {
	var rateReviewCommand commands.RateReviewCommand
	json.Unmarshal(request.Command, &rateReviewCommand)

	req := &pb.RateReviewRequest{
		ReviewId: rateReviewCommand.ReviewID,
		Star:     int32(rateReviewCommand.Star),
	}

	_, err := handler.client.RateReview(ctx, req)
	if err != nil {
		request.ErrResponse <- err
		return
	}

	fmt.Printf("Review (%d) rated with star : %d\n", rateReviewCommand.ReviewID, rateReviewCommand.Star)

	request.HandlerResponse <- fmt.Sprintf("%v", rateReviewCommand.ReviewID)
}
