package main

import (
	"context"
	"log"
	"time"
	pb "example.com/query" 
	"google.golang.org/grpc"
	"io"

	//"golang.org/x/exp/rand"
	//"gonum.org/v1/gonum/stat/distuv"
)

const (
	address = "localhost:50051"
)

func runContinuousQuery(client pb.QueryServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), 1 * time.Second)
	defer cancel()

	stream, _ := client.ContinuousQuery(ctx)

	waitc := make(chan struct{})
	go func() {
		cnt := 0
		for {
			in, err := stream.Recv()
			if err == io.EOF {
				// read done.
				close(waitc)
				log.Printf("received response cnt %v", cnt)
				return
			}
			if err != nil {
				log.Fatalf("Failed to receive a response %v", err)
			}
			cnt += 1
			if cnt % 10000 == 0 {
				//log.Printf("Got message %s at point(%d, %d)", in.Message, in.Location.Latitude, in.Location.Longitude)
				log.Printf("Got %v-th bucket: ", cnt)
				for i := 0; i < 1; i++ {
					log.Printf("%v ", in.Bucket[i])
				}
			}
		}
	}()
	for i := 0; i < 10000000; i++ {
		var q = pb.CuckooBucketQuery{BucketId: uint64(i % 10)}
		if err := stream.Send(&q); err != nil {
			log.Fatalf("Failed to send a query at %v-th package: %v", i, err)
		}
	}
	stream.CloseSend()
	<-waitc
}


func runSingleQuery(client pb.QueryServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), 1 * time.Second)
	defer cancel()

	for i := 0; i < 1000000; i++ {
		bucket, err := client.SingleQuery(ctx, &pb.CuckooBucketQuery{BucketId: uint64(i % 10)})

		if err != nil {
			log.Fatalf("failed to query %v", err)
		}
		if i % 1000 == 0 {
			log.Printf("Got %v-th bucket: ", i)
			for j := 0; j < 1; j++ {
				log.Printf("%v ", bucket.Bucket[j])
			}
		}
		/*
		log.Printf("Return Value %v", r.GetValue())

		wait_time := time.Duration(poisson_dist.Rand() * 100) * time.Millisecond;
		log.Printf("Waiting Time %v", wait_time)
		time.Sleep(wait_time)
		*/
	}
}

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Failed to connect server %v", address)
	}

	defer conn.Close()
	c := pb.NewQueryServiceClient(conn)

	//runSingleQuery(c);
	runContinuousQuery(c);

	/*
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
	*/

}