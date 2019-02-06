package actors

import (
	"fmt"
	"log"
	"time"

	"github.com/AsynkronIT/protoactor-go/actor"
)

type ReviewActor struct{}

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

		// Verify if user bought the product
		props := actor.FromProducer(NewOrdersActor)
		ordersPid := context.Spawn(props)

		future := context.RequestFuture(ordersPid, &VerifyOrderMessage{
			ProductID: msg.ProductID,
			UserID:    msg.UserID,
		}, 3*time.Second)
		result, err := future.Result()
		if err != nil {
			log.Print(err.Error())
			return
		}
		fmt.Printf("Received %#v\n", result)

		// Verify user
		usersProp := actor.FromProducer(NewUsersActor)
		usersPid := context.Spawn(usersProp)

		usersFuture := context.RequestFuture(usersPid, &VerifyUserMessage{
			UserID: msg.UserID,
		}, 3*time.Second)
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
