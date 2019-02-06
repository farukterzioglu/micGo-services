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
type OrdersActor struct {
	mpOrdersPid *actor.PID
}

// Receive ...
func (actor *OrdersActor) Receive(context actor.Context) {
	switch msg := context.Message().(type) {
	case *VerifyOrderMessage:
		fmt.Printf("VerifyOrderMessage %v\n", msg)

		actor.mpOrdersPid.
			RequestFuture(msg, 4 * time.Second).
			PipeTo(ctx.Self())
	case *VerifyMPOrderMessage:
		context.Respond(msg)
	case struct{}:
		// TODO : Get data from source
		time.Sleep(2 * time.Second)
		isAnOrder := true

		if !isAnOrder {
			context.Respond(&VerifyOrderResponse{IsPurchased: false})
			return
		}

		context.Respond(&VerifyOrderResponse{IsPurchased: true})
	}
}

// NewOrdersActor ...
func NewOrdersActor(pid *actor.PID) actor.Actor {
	return &OrdersActor{
		mpOrdersPid: pid,
	}
}
