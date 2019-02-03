package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/farukterzioglu/micGo-services/Review.API/dtos"
	pb "github.com/farukterzioglu/micGo-services/Review.CommandRpcServer/reviewservice"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

// QueryController for review queries
type QueryController struct {
	client *pb.ReviewServiceClient
}

// NewQueryController create a new QueryController instance
func NewQueryController(c *pb.ReviewServiceClient) *QueryController {
	return &QueryController{
		client: c,
	}
}

// RegisterRoutes registers routes for Review query endpoints
func (controller *QueryController) RegisterRoutes(r *mux.Router, p string) {
	ur := r.PathPrefix(p).Subrouter()

	// swagger:route GET /review QueryAPI reviewList
	// ---
	// Returns all reviews.
	//
	// responses:
	//   200: reviewsResp
	//   404: notFound
	//	 500: internal
	ur.HandleFunc("", controller.getReviews).Methods("GET")

	// swagger:route GET /review/{ReviewID} QueryAPI getReviewReq
	// ---
	// Returns a review by id.
	// If the review id is null, Error Bad Request will be returned.
	// responses:
	//   200: reviewResp
	//   404: notFound
	//	 500: internal
	ur.HandleFunc("/{ReviewID}", controller.getReview).Methods("GET")
}

func (controller *QueryController) getReview(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	reviewIDStr := params["ReviewID"]

	_, err := uuid.Parse(reviewIDStr)
	if err != nil {
		// TODO : write validation message to response
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	req := &pb.GetReviewRequest{ReviewID: reviewIDStr}
	review, err := (*controller.client).GetReview(r.Context(), req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	reviewDTO := dtos.ReviewDto{
		ID:        review.ReviewID,
		Text:      review.Text,
		Star:      int8(review.Star),
		ProductID: review.ProductID,
		UserID:    review.UserID,
	}

	if err := json.NewEncoder(w).Encode(reviewDTO); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}
	w.WriteHeader(http.StatusOK)
}

func (controller *QueryController) getReviews(w http.ResponseWriter, r *http.Request) {
	count := 10

	req := &pb.GetTopReviewsRequest{Count: int32(count)}
	stream, err := (*controller.client).GetTopReviews(r.Context(), req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	fmt.Printf("Total review count : %d\n", req.Count)

	var reviewList []dtos.ReviewDto
	for {
		review, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		reviewDTO := dtos.ReviewDto{
			Text: review.Text,
			Star: int8(review.Star),
		}
		reviewList = append(reviewList, reviewDTO)
	}

	if err := json.NewEncoder(w).Encode(reviewList); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}
	w.WriteHeader(http.StatusOK)
}
