package actors

import (
	"fmt"

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

		props := actor.FromProducer(NewOrdersActor)
		child := context.Spawn(props)
		child.Tell(VerifyOrderMessage{
			ProductID: msg.ProductID,
			UserID:    msg.UserID,
		})

		usersProp := actor.FromProducer(NewUsersActor)
		userschild := context.Spawn(usersProp)
		userschild.Tell(VerifyUserMessage{
			UserID: msg.UserID,
		})
	}
}

// NewReviewActor return a ReviewActor instance
func NewReviewActor() actor.Actor {
	return &ReviewActor{}
}
