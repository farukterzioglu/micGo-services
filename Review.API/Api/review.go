package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	_ "github.com/Shopify/sarama"
	"github.com/farukterzioglu/micGo-services/Review.API/Dtos"
	"github.com/farukterzioglu/micGo-services/Review.Domain/Commands/V1"
	"github.com/farukterzioglu/micGo-services/Review.Domain/Models"

	"github.com/gorilla/mux"
)

// ReviewRoutes for review endpoints
type ReviewRoutes struct {
	// TODO : Kafka publisher
}

// NewReviewRoutes create a new ReviewRoutes instance
// TODO : Get publisher from params
func NewReviewRoutes() *ReviewRoutes {
	return &ReviewRoutes{}
}

// RegisterReviewRoutes registers routes for Review
func (routes *ReviewRoutes) RegisterReviewRoutes(r *mux.Router, p string) {
	ur := r.PathPrefix(p).Subrouter()

	// swagger:route PUT /review CommandAPI createReviewReq
	// ---
	// summary: Creates a new review.
	// description:
	// responses:
	//   200: ok
	//   400: badReq
	ur.HandleFunc("", createReview).Methods("PUT")

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
	ur.HandleFunc("", getReviews).Methods("GET")

	// swagger:operation POST /review/{ReviewID}/ratereview CommandAPI rateReviewReq
	// ---
	// summary: Rates the review.
	// description: If the review id is null, Error Bad Request will be returned.
	// responses:
	//   "200":
	//     "$ref": "#/responses/rateReviewResp"
	//   "400":
	//     "$ref": "#/responses/badReq"
	//   "404":
	//     "$ref": "#/responses/notFound"
	//   "500":
	//     "$ref": "#/responses/internal"
	ur.HandleFunc("/{ReviewID}/ratereview", rateReview).Methods("POST")
}

func getReviews(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
	w.Write([]byte("Not implemented!"))
}

func createReview(w http.ResponseWriter, r *http.Request) {
	var review dtos.ReviewDto
	_ = json.NewDecoder(r.Body).Decode(&review)

	command := &commands.CreateReviewCommand{
		Review: models.Review{
			Text: review.Text,
			Star: review.Star,
		},
	}

	command.Review.Status = models.Created

	msg, err := json.Marshal(command)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	err = publish(string(msg), "create-review")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func rateReview(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	reviewIDStr := params["ReviewID"]

	reviewID, err := strconv.ParseInt(reviewIDStr, 10, 32)
	if err != nil {
		// TODO : write validation message to response
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var rating dtos.ReviewRatingDto
	_ = json.NewDecoder(r.Body).Decode(&rating)

	command := &commands.RateReviewCommand{
		ReviewID: (int32)(reviewID),
		Star:     rating.Star,
	}

	msg, err := json.Marshal(command)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	// TODO : Retry & circuit breake
	err = publish(string(msg), "rate-review")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func publish(message, topicName string) error {
	fmt.Printf("%s - %s\n", topicName, message)
	// msg := &sarama.ProducerMessage{
	// 	Topic: topicName,
	// 	Value: sarama.StringEncoder(message),
	// }

	// p, o, err := producer.SendMessage(msg)
	// if err != nil {
	// 	fmt.Println("Error publish: ", err.Error())
	// 	return err
	// }

	// fmt.Printf("Delivered %s[part:%d] (@%d) - %s\n'", topicName, p, o, message)
	return nil
}
