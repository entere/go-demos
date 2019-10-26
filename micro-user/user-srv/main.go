package main

import (
	"github.com/entere/go-demos/micro-user/base"
	"github.com/entere/go-demos/micro-user/user-srv/handler"
	user "github.com/entere/go-demos/micro-user/user-srv/proto/user"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/util/log"
)

func main() {
	base.Init()
	// New Service
	service := micro.NewService(
		micro.Name("io.github.entere.srv.user"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	user.RegisterUserHandler(service.Server(), new(handler.User))

	// Register Struct as Subscriber
	// micro.RegisterSubscriber("io.github.entere.srv.user", service.Server(), new(subscriber.User))

	// Register Function as Subscriber
	// micro.RegisterSubscriber("io.github.entere.srv.user", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
