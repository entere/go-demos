package main

import (
	"context"
	"net"

	pb "github.com/entere/go-demos/grpc-day01/proto"
	"google.golang.org/grpc"
)

// UserService ..
type UserService struct{}

// Info ..
func (s *UserService) Info(ctx context.Context, r *pb.UserRequest) (*pb.UserResponse, error) {

	userResponse := &pb.UserResponse{Id: r.GetId(), Name: "entere", Gender: "1", Mobile: "138xxxxxxxx"}
	return userResponse, nil
}

// PORT ..
const PORT = "9001"

func main() {
	server := grpc.NewServer()
	pb.RegisterUserServiceServer(server, &UserService{})

	lis, err := net.Listen("tcp", ":"+PORT)
	if err != nil {

	}
	server.Serve(lis)
}
