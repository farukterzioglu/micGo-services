package actors

import (
	"fmt"
	"time"

	"github.com/AsynkronIT/protoactor-go/actor"
)

// VerifyMPOrderMessage carries relation of user- marketplace product
type VerifyMPOrderMessage struct {
	ProductID string
	UserID    string
}

// VerifyMPOrderResponse ..
type VerifyMPOrderResponse struct {
	IsPurchased bool
}

// MPOrdersActor actor for order-user relation queries
type MPOrdersActor struct{}

// Receive ...
func (actor *MPOrdersActor) Receive(context actor.Context) {
	switch msg := context.Message().(type) {
	case *VerifyMPOrderMessage:
		fmt.Printf("MPOrdersActor -> VerifyMPOrderMessage %v\n", msg)

		// TODO : Get data from source
		time.Sleep(time.Second)

		// TODO : Check if it is MarketPlace order
		if msg.ProductID == "111" {
			context.Respond(&VerifyMPOrderResponse{IsPurchased: true})
			return
		}

		fmt.Printf("MPOrdersActor -> Not a marketplace order.\n")
		context.Respond(&VerifyMPOrderResponse{IsPurchased: false})
	}
}

// NewMPOrdersActor ...
func NewMPOrdersActor() actor.Actor {
	return &MPOrdersActor{}
}
