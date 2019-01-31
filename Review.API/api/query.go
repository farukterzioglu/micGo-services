package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/farukterzioglu/micGo-services/Review.API/dtos"
	pb "github.com/farukterzioglu/micGo-services/Review.CommandRpcServer/reviewservice"
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
	// summary: Gets all reviews.
	// description:
	// responses:
	//   "200":
	//     "$ref": "#/responses/reviewsResp"
	//   "404":
	//     "$ref": "#/responses/notFound"
	//   "500":
	//     "$ref": "#/responses/internal"
	ur.HandleFunc("", controller.getReviews).Methods("GET")
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

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(reviewList); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}
}
