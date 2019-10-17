package main

import (
	"context"
	"log"

	pb "github.com/entere/go-demos/grpc-day01/proto"
	"google.golang.org/grpc"
)

// PORT ..
const PORT = "9001"

func main() {
	conn, err := grpc.Dial(":"+PORT, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("grpc.Dial err: %v", err)
	}

	client := pb.NewUserServiceClient(conn)

	resp, err := client.Info(context.Background(), &pb.UserRequest{Id: "1"})
	if err != nil {
		log.Fatalf("client.User err: %v", err)
	}
	log.Printf("resp: %v", resp)
}
