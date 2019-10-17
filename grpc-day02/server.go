package main

import (
	"context"
	"log"

	pb "github.com/entere/go-demos/grpc-day02/proto"

	"github.com/micro/go-micro"
)

// User ..
type User struct {
	userService pb.UserService
}

// Info ..
func (s *User) Info(ctx context.Context, r *pb.UserRequest) (*pb.UserResponse, error) {

	userResponse := &pb.UserResponse{Id: r.GetId(), Name: "entere", Gender: "1", Mobile: "138xxxxxxxx"}
	return userResponse, nil
}

func main() {
	server := micro.NewService(
		// 必须和 consignment.proto 中的 package 一致
		micro.Name("user.service"),
		micro.Version("latest"),
	)
	server.Init()

	pb.RegisterUserServiceHandler(server.Server(), &User{})

	if err := server.Run(); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
