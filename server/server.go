package main

import (
	"context"
	"log"
	"net"
	pb "example.com/query" 
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type QueryServiceServer struct {
	pb.UnimplementedQueryServiceServer
	Database *map[int32]int32
}

func (s *QueryServiceServer) QueryDatabase(ctx context.Context, index *pb.SimpleQuery) (*pb.QueryResponse, error) {
	var id = index.GetIndex()
	log.Printf("Received: %v", id)
	var ret int32 = (*s.Database)[id]
	return &pb.QueryResponse{Value: ret}, nil;
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen %v", err)
	}

	//var db map[int32]int32
	db := make(map[int32]int32)

	for i := int32(0); i < 100; i++ {
		db[i] = i * 10
	}

	s := grpc.NewServer() 
	pb.RegisterQueryServiceServer(s, &QueryServiceServer{Database: &db})
	log.Printf("server listening at %v", lis.Addr()) 

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to server %v", err)
	}
}
