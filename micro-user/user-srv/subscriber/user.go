package subscriber

import (
	"context"

	user "github.com/entere/go-demos/micro-user/user-srv/proto/user"
	"github.com/micro/go-micro/util/log"
)

type User struct{}

func (e *User) Handle(ctx context.Context, msg *user.Message) error {
	log.Log("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *user.Message) error {
	log.Log("Function Received message: ", msg.Say)
	return nil
}
