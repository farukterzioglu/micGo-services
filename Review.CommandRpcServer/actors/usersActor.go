package actors

import (
	"fmt"

	"github.com/AsynkronIT/protoactor-go/actor"
)

// VerifyUserMessage carries verification of user
type VerifyUserMessage struct {
	UserID string
}

// UsersActor actor for order-user relation queries
type UsersActor struct{}

// Receive ...
func (actor *UsersActor) Receive(context actor.Context) {
	switch msg := context.Message().(type) {
	case VerifyUserMessage:
		fmt.Printf("VerifyUserMessage %v\n", msg)
	}
}

// NewUsersActor return a new UsersActor instance
func NewUsersActor() actor.Actor {
	return &UsersActor{}
}
