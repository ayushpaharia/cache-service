package main

import (
	"context"
	"log"
	"net"

	pb "cache-service/proto"
	"cache-service/store"

	"google.golang.org/grpc"
)

const (
	port = ":5001"
)

var DB store.KVStore

type RPCServer struct {
	pb.UnimplementedCacheServiceServer
}

// Get Method response
func (s *RPCServer) Get(ctx context.Context, in *pb.GetRequest) (*pb.ServerResponse, error) {
	value, err := DB.Get(in.GetKey())
	if err != nil {
		return &pb.ServerResponse{Success: false, Value: "", Error: err.Error()}, nil
	}
	log.Printf("[GET]%v", value)
	return &pb.ServerResponse{Success: true, Value: value, Error: ""}, nil
}

// Set Method response
func (s *RPCServer) Set(ctx context.Context, in *pb.SetRequest) (*pb.ServerResponse, error) {
	_, err := DB.Set(in.GetKey(), in.GetValue())
	if err != nil {
		return &pb.ServerResponse{Success: false, Value: "", Error: err.Error()}, nil
	}
	log.Printf("[SET]%s:%s", in.GetKey(), in.GetValue())
	return &pb.ServerResponse{Success: true, Value: in.GetValue(), Error: ""}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()

	DB, err = store.CreateKVStore()
	if err != nil {
		log.Fatalf("Failed to create DB: %v", err)
		panic(err)
	}

	pb.RegisterCacheServiceServer(grpcServer, &RPCServer{})
	log.Printf("server listening at %v", lis.Addr())

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
