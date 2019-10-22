package main

import (
	"github.com/entere/go-demos/micro-part02/user-srv/handler"
	"github.com/entere/go-demos/micro-part02/user-srv/subscriber"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/consul"
	"github.com/micro/go-micro/util/log"

	user "github.com/entere/go-demos/micro-part02/user-srv/proto/user"
)

func main() {
	// New Registry 使用consul注册
	reg := consul.NewRegistry(func(op *registry.Options) {
		op.Addrs = []string{
			"127.0.0.1:8500",
		}
	})
	// New Service
	service := micro.NewService(
		micro.Registry(reg),
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
