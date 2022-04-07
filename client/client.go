package main

import (
	"context"
	"log"
	"time"
	pb "example.com/query" 
	"google.golang.org/grpc"

	"golang.org/x/exp/rand"
	"gonum.org/v1/gonum/stat/distuv"
)

const (
	address = "localhost:50051"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Failed to connect server %v", address)
	}

	defer conn.Close()
	c := pb.NewQueryServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 20 * time.Second)
	defer cancel()

	poisson_dist := distuv.Poisson{Lambda: 2, Src: rand.NewSource(100)}
	log.Printf("Mean of Poisson %v", poisson_dist.Mean())

	for i := 0; i < 10; i++ {
		r, err := c.QueryDatabase(ctx, &pb.SimpleQuery{Index: int32(i)})

		if err != nil {
			log.Fatalf("failed to query %v", err)
		}

		log.Printf("Return Value %v", r.GetValue())

		wait_time := time.Duration(poisson_dist.Rand() * 100) * time.Millisecond;
		log.Printf("Waiting Time %v", wait_time)
		time.Sleep(wait_time)
	}

}