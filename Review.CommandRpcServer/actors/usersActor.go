package actors

import (
	"fmt"
	"time"

	"github.com/AsynkronIT/protoactor-go/actor"
)

// VerifyUserMessage hold user id for verification
type VerifyUserMessage struct {
	UserID string
}

// VerifyUserResponse presents if user is permitted to comment
type VerifyUserResponse struct {
	IsPermitted bool
}

// UsersActor actor for order-user relation queries
type UsersActor struct{}

// Receive ...
func (actor *UsersActor) Receive(context actor.Context) {
	switch msg := context.Message().(type) {
	case *VerifyUserMessage:
		fmt.Printf("VerifyUserMessage %v\n", msg)

		// TODO : Get data from source
		time.Sleep(1 * time.Second)
		context.Respond(&VerifyUserResponse{IsPermitted: true})
	}
}

// NewUsersActor return a new UsersActor instance
func NewUsersActor() actor.Actor {
	return &UsersActor{}
}
