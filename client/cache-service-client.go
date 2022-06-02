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
	address = "localhost:50051"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()
	c := pb.NewCacheServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	var kvstore = make(map[string]string)
	kvstore["key1"] = "value1"
	kvstore["key2"] = "value2"

	// for key, val := range kvstore {
	// 	r, err := c.Get(ctx, &pb.GetRequest{Key: key})
	// 	if err != nil {
	// 		log.Fatalf("could not set: %v", err)
	// 	}
	// 	log.Printf("Key: %s, Value: %s", r.GetValue(), val)
	// }
	for key, val := range kvstore {
		r, err := c.Set(ctx, &pb.SetRequest{Key: key, Value: val})
		if err != nil {
			log.Fatalf("could not set: %v", err)
		}
		log.Printf("Key: %s, Value: %s", r.GetValue(), val)
	}

	// r, err := c.Get(ctx, &pb.GetRequest{Key: "gRPC"})
	// if err != nil {
	// 	log.Fatalf("could not greet: %v", err)
	// }
	// log.Printf("Greeting: %s", r.Value)
}
