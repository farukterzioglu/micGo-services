package actors

import (
	"fmt"
	"log"
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
		fmt.Printf("OrdersActor -> VerifyOrderMessage %v\n", msg)

		// Verify from marketplace
		request := &VerifyMPOrderMessage{
			ProductID: msg.ProductID,
			UserID:    msg.UserID,
		}
		future := actor.mpOrdersPid.RequestFuture(request, 4*time.Second)

		futureResult, err := future.Result()
		if err != nil {
			log.Print(err.Error())
			return
		}

		var mpResult *VerifyMPOrderResponse
		mpResult = futureResult.(*VerifyMPOrderResponse)

		if mpResult.IsPurchased {
			fmt.Printf("OrdersActor -> Verified from market place.\n")
			context.Respond(&VerifyOrderResponse{IsPurchased: true})
			return
		}

		// Verify non-marketplace product
		// TODO : Get data from source
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
