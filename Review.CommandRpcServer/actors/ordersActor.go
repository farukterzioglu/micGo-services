package actors

import (
	"fmt"
	"time"

	"github.com/AsynkronIT/protoactor-go/actor"
)

// VerifyOrderMessage carries relation of user-product
type VerifyOrderMessage struct {
	ProductID string
	UserID    string
}

// VerifyOrderResponse ..
type VerifyOrderResponse struct {
	IsPurchased bool
}

// OrdersActor actor for order-user relation queries
type OrdersActor struct{}

// Receive ...
func (actor *OrdersActor) Receive(context actor.Context) {
	switch msg := context.Message().(type) {
	case *VerifyOrderMessage:
		fmt.Printf("VerifyOrderMessage %v\n", msg)

		// TODO : Get data from source
		time.Sleep(2 * time.Second)
		context.Respond(&VerifyOrderResponse{IsPurchased: true})
	}
}

// NewOrdersActor ...
func NewOrdersActor() actor.Actor {
	return &OrdersActor{}
}
