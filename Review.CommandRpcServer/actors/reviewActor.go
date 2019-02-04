package actors

import (
	"fmt"

	"github.com/AsynkronIT/protoactor-go/actor"
)

type ReviewActor struct {
	// ordersPid *actor.PID
}

type SaveReviewMessage struct {
	ID        string
	ProductID string
	UserID    string
	Text      string
	Star      int8
}

func (actor *ReviewActor) Receive(context actor.Context) {

	// TODO : Send to Orders actor to check for buyer
	// TODO : Send to Users actor for review approval

	switch msg := context.Message().(type) {
	case SaveReviewMessage:
		fmt.Printf("Message %v\n", msg)
	}
}

func NewReviewActor() actor.Actor {
	return &ReviewActor{}
}
