package main

import (
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-micro"
	"micro-day01/user-srv/handler"
	"micro-day01/user-srv/subscriber"

	user "micro-day01/user-srv/proto/user"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("mu.micro.ci.srv.user"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	user.RegisterUserHandler(service.Server(), new(handler.User))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("mu.micro.ci.srv.user", service.Server(), new(subscriber.User))

	// Register Function as Subscriber
	micro.RegisterSubscriber("mu.micro.ci.srv.user", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
