package actors

import (
	"fmt"
	"log"
	"time"

	"github.com/AsynkronIT/protoactor-go/actor"
)

// ReviewActor ...
type ReviewActor struct{}

// SaveReviewMessage ...
type SaveReviewMessage struct {
	ID        string
	ProductID string
	UserID    string
	Text      string
	Star      int8
}

// Receive handles SaveReviewMessage, ... messages
func (reviewActor *ReviewActor) Receive(context actor.Context) {

	switch msg := context.Message().(type) {
	case SaveReviewMessage:
		fmt.Printf("SaveReviewMessage %v\n", msg)

		// Create actors
		mpOrdersProps := actor.FromProducer(NewMPOrdersActor)
		mpOrdersPid := context.Spawn(mpOrdersProps)

		props := actor.FromFunc(func(ctx actor.Context) {
			return NewOrdersActor(mpOrdersPid)
		})
		ordersPid := context.Spawn(props)

		usersProp := actor.FromProducer(NewUsersActor)
		usersPid := context.Spawn(usersProp)

		//// Send requests
		// Verify if user bought the product
		future := context.RequestFuture(ordersPid, &VerifyOrderMessage{
			ProductID: msg.ProductID,
			UserID:    msg.UserID,
		}, 3*time.Second)

		// Verify user
		usersFuture := context.RequestFuture(usersPid, &VerifyUserMessage{
			UserID: msg.UserID,
		}, 3*time.Second)

		//// Get results
		// Get verify order result
		result, err := future.Result()
		if err != nil {
			log.Print(err.Error())
			return
		}
		fmt.Printf("Received %#v\n", result)

		// Get verify user result
		usersresult, err := usersFuture.Result()
		if err != nil {
			log.Print(err.Error())
			return
		}
		fmt.Printf("Received %#v\n", usersresult)
	}
}

// NewReviewActor return a ReviewActor instance
func NewReviewActor() actor.Actor {
	return &ReviewActor{}
}
