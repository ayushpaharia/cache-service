package main

import (
	"context"
	"log"
	"time"

	pb "cache-service/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	address = "localhost:5001"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	log.Println("Connected to server")

	defer conn.Close()
	c := pb.NewCacheServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.Set(ctx, &pb.SetRequest{Key: "key", Value: "value"})
	if err != nil {
		log.Fatalf("could not set: %v", err)
	}
	log.Printf("Key: %s, Value: %s", r.GetValue(), r.GetValue())

}
