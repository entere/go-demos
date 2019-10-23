package main

import (
	"github.com/entere/go-demos/micro-file-upload/upload-srv/handler"
	"github.com/entere/go-demos/micro-file-upload/upload-srv/subscriber"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/util/log"

	upload "github.com/entere/go-demos/micro-file-upload/upload-srv/proto/upload"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("io.github.entere.srv.upload"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	upload.RegisterUploadHandler(service.Server(), new(handler.Upload))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("io.github.entere.srv.upload", service.Server(), new(subscriber.Upload))

	// Register Function as Subscriber
	micro.RegisterSubscriber("io.github.entere.srv.upload", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
