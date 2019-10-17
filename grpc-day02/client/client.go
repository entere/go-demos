package main

import (
	"context"
	"log"

	pb "github.com/entere/go-demos/grpc-day02/proto"
	microclient "github.com/micro/go-micro/client"
	"github.com/micro/go-micro/config/cmd"
)

// PORT ..
const PORT = "9001"

func main() {
	cmd.Init()

	client := pb.NewUserServiceClient("go.micro.srv.user", microclient.DefaultClient)

	resp, err := client.Info(context.Background(), &pb.UserRequest{Id: "1"})
	if err != nil {
		log.Fatalf("client.User err: %v", err)
	}
	log.Printf("resp: %v", resp)
}
