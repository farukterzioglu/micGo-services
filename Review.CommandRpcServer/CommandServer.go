package main

import (
	"context"
	"io"
	"log"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	"github.com/AsynkronIT/protoactor-go/actor"
	"github.com/farukterzioglu/micGo-services/Review.CommandRpcServer/actors"
	pb "github.com/farukterzioglu/micGo-services/Review.CommandRpcServer/reviewservice"
	"github.com/farukterzioglu/micGo-services/Review.Domain/Models"
)

// CommandServer for handling rpc commands
type CommandServer struct{}

// NewCommandServer creates and return a CommandServer instance
func NewCommandServer() *CommandServer {
	s := &CommandServer{}
	return s
}

// SaveReview handles SaveReview rpc command
func (server *CommandServer) SaveReview(ctx context.Context, request *pb.NewReviewRequest) (*pb.ReviewId, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		requestOwner := md["request-owner"]
		log.Printf("SaveReview request owner %v\n", requestOwner)
	}

	chn := make(chan interface{})

	props := actor.FromProducer(func() actor.Actor {
		reviewActor := actors.ReviewActor{
			ID:          request.Review.ReviewID,
			ProductID:   request.Review.ProductID,
			UserID:      request.Review.UserID,
			Text:        request.Review.Text,
			Star:        int8(request.Review.Star),
			ResponseChn: chn,
		}
		return &reviewActor
	})
	reviewsPid := actor.Spawn(props)
	reviewsPid.Tell(&actors.SaveReviewMessage{})

	select {
	case result := <-chn:
		log.Printf("CommandServer -> SaveReview response %v\n", result)
		return &pb.ReviewId{ReviewId: request.Review.ReviewID}, nil
	case <-time.After(5 * time.Second):
		log.Printf("CommandServer -> SaveReview request timed out!\n")
		return nil, status.Errorf(codes.Canceled, "request timed out!")
	}
}

// SaveReviews handles SaveReviews rpc command
func (server *CommandServer) SaveReviews(stream pb.ReviewService_SaveReviewsServer) error {
	md, _ := metadata.FromIncomingContext(stream.Context())
	_ = md["batchCount"] // batchCount

	for {
		request, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		var review models.Review
		review = models.Review{
			Text: request.Review.Text,
			Star: int8(request.Review.Star),
		}
		log.Printf("Received review with text : %s\n", review.Text)

		var reviewID string
		// TODO : Process review
		reviewID = "0000"

		if err := stream.Send(&pb.ReviewId{ReviewId: reviewID}); err != nil {
			return err
		}
	}
}

// GetReview returns a review by id=req.ReviewID
func (server *CommandServer) GetReview(ctx context.Context, req *pb.GetReviewRequest) (*pb.Review, error) {
	reviewID := req.ReviewID

	// TODO : Query data source
	return &pb.Review{
		Star:     1,
		Text:     "Sample review",
		ReviewID: reviewID,
	}, nil
}

// GetTopReviews returns top 'GetTopReviewsRequest.count' reviews
func (server *CommandServer) GetTopReviews(req *pb.GetTopReviewsRequest, stream pb.ReviewService_GetTopReviewsServer) error {
	// TODO : Get reviews
	reviewList := []models.Review{
		models.Review{
			Text: "First awesome review",
			Star: 5,
		},
		models.Review{
			Text: "Last meh review",
			Star: 3,
		},
	}

	for _, review := range reviewList {
		reviewReq := pb.Review{
			Text: review.Text,
			Star: int32(review.Star),
		}

		if err := stream.Send(&reviewReq); err != nil {
			return err
		}
	}

	return nil
}

// RateReview saves the rating for review
func (server *CommandServer) RateReview(ctx context.Context, req *pb.RateReviewRequest) (*pb.Empty, error) {
	// TODO : Save the rating
	log.Printf("Rated -> review id : %s, rating : %d\n", req.ReviewId, req.Star)

	return &pb.Empty{}, nil
}
