package subscriber

import (
	"context"

	"github.com/micro/go-micro/util/log"

	upload "github.com/entere/go-demos/micro-file-upload/upload-srv/proto/upload"
)

type Upload struct{}

func (e *Upload) Handle(ctx context.Context, msg *upload.Message) error {
	log.Log("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *upload.Message) error {
	log.Log("Function Received message: ", msg.Say)
	return nil
}
