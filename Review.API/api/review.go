package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Shopify/sarama"
	"github.com/farukterzioglu/micGo-services/Review.API/dtos"
	"github.com/farukterzioglu/micGo-services/Review.Domain/Commands/V1"
	"github.com/farukterzioglu/micGo-services/Review.Domain/Models"
	"github.com/google/uuid"

	"github.com/gorilla/mux"
)

var _topicName = "review-commands"

// ReviewRoutes for review endpoints
type ReviewRoutes struct {
	producer *sarama.SyncProducer
}

// NewReviewRoutes create a new ReviewRoutes instance
func NewReviewRoutes(p *sarama.SyncProducer) *ReviewRoutes {
	return &ReviewRoutes{
		producer: p,
	}
}

// RegisterReviewRoutes registers routes for Review
func (routes *ReviewRoutes) RegisterReviewRoutes(r *mux.Router, p string) {
	ur := r.PathPrefix(p).Subrouter()

	// swagger:route PUT /review CommandAPI createReviewReq
	// ---
	// Creates a new review.
	// Creates a 'create review commnand' and sends to Kafka
	// responses:
	//   202: ok
	//   400: badReq
	ur.HandleFunc("", routes.createReview).Methods("PUT")

	// swagger:operation POST /review/{ReviewID}/ratereview CommandAPI rateReviewReq
	// ---
	// summary: Rates the review.
	// description: If the review id is null, Error Bad Request will be returned.
	// responses:
	//   "202":
	//     "$ref": "#/responses/rateReviewResp"
	//   "400":
	//     "$ref": "#/responses/badReq"
	//   "404":
	//     "$ref": "#/responses/notFound"
	//   "500":
	//     "$ref": "#/responses/internal"
	ur.HandleFunc("/{ReviewID}/ratereview", routes.rateReview).Methods("POST")
}

func (routes *ReviewRoutes) createReview(w http.ResponseWriter, r *http.Request) {
	var review dtos.CreateReviewDto
	_ = json.NewDecoder(r.Body).Decode(&review)

	id, _ := uuid.NewRandom()
	reviewID := id.String()
	command := &commands.CreateReviewCommand{
		Review: models.Review{
			ID:        reviewID,
			ProductID: review.ProductID,
			Text:      review.Text,
			Star:      review.Star,
			UserID:    review.UserID,
		},
	}

	command.Review.Status = models.Created

	msg, err := json.Marshal(command)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	cmdMessage := models.CommandMessage{
		CommandData: msg,
		CommandType: "create-review",
	}
	cmdMessageStr, err := json.Marshal(cmdMessage)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	err = publish(routes.producer, cmdMessageStr, reviewID, _topicName)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	// TODO : Gives a warning : 'http: multiple response.WriteHeader calls'
	w.Header().Set("Content-Location", "/review/"+reviewID)
	json.NewEncoder(w).Encode(reviewID)
	w.WriteHeader(http.StatusAccepted)
}

func (routes *ReviewRoutes) rateReview(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	reviewIDStr := params["ReviewID"]

	reviewID, err := uuid.Parse(reviewIDStr)
	if err != nil {
		// TODO : write validation message to response
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var rating dtos.ReviewRatingDto
	_ = json.NewDecoder(r.Body).Decode(&rating)

	command := &commands.RateReviewCommand{
		ReviewID: reviewID.String(),
		Star:     rating.Star,
	}

	msg, err := json.Marshal(command)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	cmdMessage := models.CommandMessage{
		CommandData: msg,
		CommandType: "rate-review",
	}
	cmdMessageStr, err := json.Marshal(cmdMessage)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	// TODO : Retry & circuit breake
	err = publish(routes.producer, cmdMessageStr, reviewIDStr, _topicName)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusAccepted)
}

func publish(producer *sarama.SyncProducer, message []byte, key, topicName string) error {
	// TODO : Use byte encoder
	value := string(message)

	msg := &sarama.ProducerMessage{
		Topic: topicName,
		Value: sarama.StringEncoder(value),
		Key:   sarama.StringEncoder(key),
	}

	p, o, err := (*producer).SendMessage(msg)
	if err != nil {
		fmt.Println("Error publish: ", err.Error())
		return err
	}

	log.Printf("Delivered %s[part:%d] (@%d) (key:%s) - %s\n", topicName, p, o, msg.Key, value)
	return nil
}
