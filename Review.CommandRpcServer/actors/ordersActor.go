package actors

import (
	"fmt"

	"github.com/AsynkronIT/protoactor-go/actor"
)

// VerifyOrderMessage carries relation of user-product
type VerifyOrderMessage struct {
	ProductID string
	UserID    string
}

// OrdersActor actor for order-user relation queries
type OrdersActor struct{}

func (actor *OrdersActor) Receive(context actor.Context) {
	switch msg := context.Message().(type) {
	case VerifyOrderMessage:
		fmt.Printf("VerifyOrderMessage %v\n", msg)
	}
}

func NewOrdersActor() actor.Actor {
	return &OrdersActor{}
}
