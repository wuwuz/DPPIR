package main

import (
	"context"
	"log"
	"net"
	pb "example.com/query" 
	"google.golang.org/grpc"
	"example.com/cuckoo"
	"io"
)

const (
	port = ":50051"
)

type QueryServiceServer struct {
	pb.UnimplementedQueryServiceServer
	//Database *map[int32]int32
	HashTable *cuckoo.Cuckoo 
}

func (s *QueryServiceServer) SingleQuery(ctx context.Context, in *pb.CuckooBucketQuery) (*pb.CuckooBucketResponse, error) {
	bucketId := in.GetBucketId()
	bucket := s.HashTable.GetBucketCopy(bucketId)

	//var id = index.GetIndex()
	//log.Printf("Received: %v", id)
	//var ret int32 = (*s.Database)[id]
	var ret = pb.CuckooBucketResponse{Bucket: bucket.Slot[0:4]}
	return &ret, nil;
	//return &pb.QueryResponse{Value: ret}, nil;
}

func (s *QueryServiceServer) ContinuousQuery(stream pb.QueryService_ContinuousQueryServer) error {
	for {
		request, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		bucketId := request.GetBucketId()
		bucket := s.HashTable.GetBucketCopy(bucketId)
		var ret = pb.CuckooBucketResponse{Bucket: bucket.Slot[0:4]}

		if err := stream.Send(&ret); err != nil {
			return err
		}
	}
}

func (s *QueryServiceServer) GetHashTableInfo(ctx context.Context, in *pb.HashTableInfoQuery) (*pb.HashTableInfoResponse, error) {
	var ret = pb.HashTableInfoResponse{Size: s.HashTable.Size, Load: s.HashTable.Load}
	return &ret, nil;
	//return &pb.QueryResponse{Value: ret}, nil;
}



func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen %v", err)
	}

	//var db map[int32]int32
	//db := make(map[int32]int32)

	//for i := int32(0); i < 100; i++ {
	//	db[i] = i * 10
//	}

	ht := cuckoo.NewCuckooHashTableGivenLogSize(14)
	for i := 1; i < 100; i++ {
		if ht.Insert(uint64(i)) == false {
			log.Printf("Failed to insert cuckoo hash table at %v", i) 
		}
	}

	s := grpc.NewServer() 
	pb.RegisterQueryServiceServer(s, &QueryServiceServer{HashTable: ht})
	log.Printf("server listening at %v", lis.Addr()) 

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to server %v", err)
	}
}
