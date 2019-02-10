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
		fmt.Printf("OrdersActor -> VerifyOrderMessage %v\n", msg)

		request := &VerifyMPOrderMessage{
			ProductID: msg.ProductID,
			UserID:    msg.UserID,
		}
		actor.mpOrdersPid.Request(request, context.Self())
	case *NotVerifiedByMarketPlace:
		fmt.Printf("OrdersActor -> NotVerifiedByMarketPlace %v\n", msg)

		// Verify non-marketplace product
		// TODO : Get data from source
		isAnOrder := true

		if isAnOrder {
			context.Respond(&VerifyOrderResponse{IsPurchased: true})
			return
		}
		context.Respond(&VerifyOrderResponse{IsPurchased: false})
	case *VerifiedByMarketPlace:
		fmt.Printf("OrdersActor -> VerifiedByMarketPlace %v\n", msg)
		context.Respond(&VerifyOrderResponse{IsPurchased: true})
	}
}

// NewOrdersActor ...
func NewOrdersActor(pid *actor.PID) actor.Actor {
	return &OrdersActor{mpOrdersPid: pid}
}
