package main

import (
	"bufio"
	"context"
	"io"
	"log"
	"net"
	"os"
	"strconv"
	"strings"

	"example.com/cuckoo"
	pb "example.com/query"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

var eps float64
var delta float64
var listSize uint64
var realQueryNum uint64

type QueryServiceServer struct {
	pb.UnimplementedQueryServiceServer
	//Database *map[int32]int32
	HashTable *cuckoo.Cuckoo
}

func (s *QueryServiceServer) HandleSingleQuery(in *pb.CuckooBucketQuery) *pb.CuckooBucketResponse {
	bucketId := in.GetBucketId()
	uniqueId := in.GetUniqueId()
	oneTimePad_0 := in.GetOneTimePad_0()
	oneTimePad_1 := in.GetOneTimePad_1()
	oneTimePad_2 := in.GetOneTimePad_2()
	oneTimePad_3 := in.GetOneTimePad_3()
	bucket := s.HashTable.GetBucketCopy(bucketId)

	/*
		for i := 0; i < cuckoo.BucketSize; i++ {
			bucket.Slot[i] = bucket.Slot[i] ^ oneTimePad[i]
		}
	*/

	return &pb.CuckooBucketResponse{
		UniqueId: uniqueId,
		Bucket_0: bucket.Slot[0] ^ oneTimePad_0,
		Bucket_1: bucket.Slot[1] ^ oneTimePad_1,
		Bucket_2: bucket.Slot[2] ^ oneTimePad_2,
		Bucket_3: bucket.Slot[3] ^ oneTimePad_3,
	}
}

func (s *QueryServiceServer) HandleBatchedQuery(in *pb.BatchedCuckooBucketQuery) *pb.BatchedCuckooBucketResponse {
	num := in.GetQueryNum()
	batchedQuery := in.GetBatchedQuery()
	batchedResponse := make([]*pb.CuckooBucketResponse, num)
	for i := uint64(0); i < num; i++ {
		batchedResponse[i] = s.HandleSingleQuery(batchedQuery[i])
	}

	return &pb.BatchedCuckooBucketResponse{ResponseNum: num, BatchedResponse: batchedResponse}
}

func (s *QueryServiceServer) SingleQuery(ctx context.Context, in *pb.CuckooBucketQuery) (*pb.CuckooBucketResponse, error) {
	return s.HandleSingleQuery(in), nil
	//return &pb.QueryResponse{Value: ret}, nil;
}

func (s *QueryServiceServer) ContinuousQuery(stream pb.QueryService_ContinuousQueryServer) error {
	for {
		//in, err := stream.Recv()
		in, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		ret := s.HandleBatchedQuery(in)

		if err := stream.Send(ret); err != nil {
			return err
		}
	}
}

func (s *QueryServiceServer) GetHashTableInfo(ctx context.Context, in *pb.HashTableInfoQuery) (*pb.HashTableInfoResponse, error) {
	var ret = pb.HashTableInfoResponse{Size: s.HashTable.Size, Load: s.HashTable.Load}
	return &ret, nil
	//return &pb.QueryResponse{Value: ret}, nil;
}

func readConfigInfo() (float64, float64, uint64, uint64) {
	file, err := os.Open("/root/DPPIR/config.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	line, _, err := reader.ReadLine()
	if err != nil {
		log.Fatal(err)
	}
	split := strings.Split(string(line), " ")
	var eps float64
	var delta float64
	var listSize int64
	var queryNum int64

	if eps, err = strconv.ParseFloat(split[0], 64); err != nil {
		log.Fatal(err)
	}
	if delta, err = strconv.ParseFloat(split[1], 64); err != nil {
		log.Fatal(err)
	}
	if listSize, err = strconv.ParseInt(split[2], 10, 64); err != nil {
		log.Fatal(err)
	}
	if queryNum, err = strconv.ParseInt(split[3], 10, 64); err != nil {
		log.Fatal(err)
	}

	log.Printf("%v %v %v %v", eps, delta, listSize, queryNum)

	return eps, delta, uint64(listSize), uint64(queryNum)
}

func main() {
	eps, delta, listSize, realQueryNum = readConfigInfo()

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen %v", err)
	}

	//var db map[int32]int32
	//db := make(map[int32]int32)

	//for i := int32(0); i < 100; i++ {
	//	db[i] = i * 10
	//	}

	//ht := cuckoo.NewCuckooHashTableGivenLogSize(15)
	ht := cuckoo.NewCuckooHashTable(listSize)
	for i := uint64(1); i <= listSize; i++ {
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
