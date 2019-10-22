package subscriber

import (
	"context"

	"github.com/micro/go-micro/util/log"

	upload "github.com/entere/go-demos/micro-upload/upload-srv/proto/upload"
)

// Upload ...
type Upload struct{}

// Handle ...
func (e *Upload) Handle(ctx context.Context, msg *upload.Message) error {
	log.Log("Handler Received message: ", msg.Say)
	return nil
}

// Handler ...
func Handler(ctx context.Context, msg *upload.Message) error {
	log.Log("Function Received message: ", msg.Say)
	return nil
}
