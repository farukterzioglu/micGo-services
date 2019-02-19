package actors

import (
	"log"

	"github.com/AsynkronIT/protoactor-go/actor"
)

// ReviewActor ...
type ReviewActor struct {
	ID          string
	ProductID   string
	UserID      string
	Text        string
	Star        int8
	ResponseChn chan<- interface{}

	isVerified     bool
	isUserVerified bool
}

// SaveReviewMessage ...
type SaveReviewMessage struct{}

// Receive handles SaveReviewMessage, ... messages
func (reviewActor *ReviewActor) Receive(context actor.Context) {
	switch msg := context.Message().(type) {
	case *SaveReviewMessage:
		log.Printf("ReviewActor -> SaveReviewMessage %v\n", msg)

		mpOrdersProps := actor.FromProducer(NewMPOrdersActor)
		mpOrdersPid := context.Spawn(mpOrdersProps)

		props := actor.FromProducer(func() actor.Actor {
			return NewOrdersActor(mpOrdersPid, context.Self())
		})
		ordersPid := context.Spawn(props)

		context.Request(ordersPid, &VerifyOrderMessage{
			ProductID: reviewActor.ProductID,
			UserID:    reviewActor.UserID,
		})
	case *VerifyOrderResponse:
		log.Printf("ReviewActor -> VerifyOrderResponse %v\n", msg)
		reviewActor.isVerified = msg.IsPurchased

		usersProp := actor.FromProducer(NewUsersActor)
		usersPid := context.Spawn(usersProp)

		context.Request(usersPid, &VerifyUserMessage{
			UserID: reviewActor.UserID,
		})
	case *VerifyUserResponse:
		log.Printf("ReviewActor -> VerifyUserResponse %v\n", msg)
		reviewActor.isUserVerified = msg.IsPermitted

		reviewActor.ResponseChn <- struct {
			IsVerified     bool
			IsUserVerified bool
		}{
			IsVerified:     reviewActor.isVerified,
			IsUserVerified: reviewActor.isUserVerified,
		}
	}
}
