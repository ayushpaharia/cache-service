package main

import (
	"context"
	"log"
	"net"

	pb "cache-service/proto"

	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type RPCServer struct {
	pb.UnimplementedCacheServiceServer
}

// Get Method response
func (s *RPCServer) Get(ctx context.Context, in *pb.GetRequest) (*pb.ServerResponse, error) {
	log.Printf("Get: %v", in.GetKey())
	return &pb.ServerResponse{Success: true, Value: in.GetKey(), Error: ""}, nil
}

// Set Method response
func (s *RPCServer) Set(ctx context.Context, in *pb.SetRequest) (*pb.ServerResponse, error) {
	log.Printf("Set: %s:[%s]", in.GetKey(), in.GetValue())
	return &pb.ServerResponse{Success: true, Value: in.GetValue(), Error: ""}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()

	pb.RegisterCacheServiceServer(grpcServer, &RPCServer{})
	log.Printf("server listening at %v", lis.Addr())

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
