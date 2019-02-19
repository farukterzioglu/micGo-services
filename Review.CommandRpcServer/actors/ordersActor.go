package actors

import (
	"log"

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
	reviewPid   *actor.PID
	mpOrdersPid *actor.PID
}

// Receive ...
func (actor *OrdersActor) Receive(context actor.Context) {
	switch msg := context.Message().(type) {
	case *VerifyOrderMessage:
		log.Printf("OrdersActor -> VerifyOrderMessage %v\n", msg)

		request := &VerifyMPOrderMessage{
			ProductID: msg.ProductID,
			UserID:    msg.UserID,
		}
		actor.mpOrdersPid.Request(request, context.Self())
	case *NotVerifiedByMarketPlace:
		log.Printf("OrdersActor -> NotVerifiedByMarketPlace %v\n", msg)

		// Verify non-marketplace product
		// TODO : How to propagate error
		isAnOrder, _ := verifyOrder(msg.ProductID, msg.UserID)

		if isAnOrder {
			actor.reviewPid.Tell(&VerifyOrderResponse{IsPurchased: true})
			return
		}
		actor.reviewPid.Tell(&VerifyOrderResponse{IsPurchased: false})
	case *VerifiedByMarketPlace:
		log.Printf("OrdersActor -> VerifiedByMarketPlace %v\n", msg)
		actor.reviewPid.Tell(&VerifyOrderResponse{IsPurchased: true})
	}
}

func verifyOrder(productID, userID string) (bool, error) {
	// TODO : Get data from source
	return true, nil
}

// NewOrdersActor ...
func NewOrdersActor(pid *actor.PID, rPid *actor.PID) actor.Actor {
	return &OrdersActor{mpOrdersPid: pid, reviewPid: rPid}
}
