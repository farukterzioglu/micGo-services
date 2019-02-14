package actors

import (
	"log"
	"time"

	"github.com/AsynkronIT/protoactor-go/actor"
)

// VerifyMPOrderMessage carries relation of user- marketplace product
type VerifyMPOrderMessage struct {
	ProductID string
	UserID    string
}

// NotVerifiedByMarketPlace ...
type NotVerifiedByMarketPlace struct {
	ProductID string
	UserID    string
}

// VerifiedByMarketPlace ...
type VerifiedByMarketPlace struct{}

// MPOrdersActor actor for order-user relation queries
type MPOrdersActor struct{}

// Receive ...
func (actor *MPOrdersActor) Receive(context actor.Context) {
	switch msg := context.Message().(type) {
	case *VerifyMPOrderMessage:
		log.Printf("MPOrdersActor -> VerifyMPOrderMessage %v\n", msg)

		// TODO : Get data from source
		time.Sleep(time.Second)

		// TODO : Check if it is MarketPlace order
		if msg.ProductID == "111" {
			context.Sender().Tell(&VerifiedByMarketPlace{})
			return
		}

		log.Printf("MPOrdersActor -> Not a marketplace order.\n")
		context.Sender().Tell(&NotVerifiedByMarketPlace{})
	}
}

// NewMPOrdersActor ...
func NewMPOrdersActor() actor.Actor {
	return &MPOrdersActor{}
}
