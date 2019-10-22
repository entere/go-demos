package main

import (
	"github.com/entere/go-demos/micro-upload/upload-srv/handler"
	"github.com/entere/go-demos/micro-upload/upload-srv/subscriber"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/util/log"

	upload "github.com/entere/go-demos/micro-upload/upload-srv/proto/upload"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.ci.srv.upload"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	upload.RegisterUploadHandler(service.Server(), new(handler.Upload))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.micro.ci.srv.upload", service.Server(), new(subscriber.Upload))

	// Register Function as Subscriber
	micro.RegisterSubscriber("go.micro.ci.srv.upload", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
