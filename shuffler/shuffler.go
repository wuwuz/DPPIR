package main

import (
	"context"
	"log"
	"time"
	pb "example.com/query" 
	"google.golang.org/grpc"
	"io"

	"golang.org/x/exp/rand"
	"gonum.org/v1/gonum/stat/distuv"

	"example.com/cuckoo"
)

const (
	address = "localhost:50051"
	eps = 0.5
	delta = 1e-6
)

var paddingNum int
var hashTableSize uint64

func calcLaplacePadding(eps float64, delta float64) int {
	LaplaceDist := distuv.Laplace{Mu: 0, Scale: 1.0 / eps, Src: rand.NewSource(100)}
	for i := 1; i <= 1000; i++ {
		if LaplaceDist.Survival(float64(i)) < delta {
			return i
		}
	}
	return 1000
}

func runContinuousQuery(client pb.QueryServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), 2 * time.Second)
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

	//timer set
	// generate some real queries and put them into a hash table
	queryBuffer := make([]pb.CuckooBucketQuery, 0, 200000)
	for i := 0; i < 10000; i++ {
		h1, h2 := cuckoo.GetTwoHash(uint64(i));
		h1 = h1 % hashTableSize;
		h2 = h2 % hashTableSize;

		queryBuffer = append(queryBuffer, pb.CuckooBucketQuery{BucketId: h1})
		queryBuffer = append(queryBuffer, pb.CuckooBucketQuery{BucketId: h2})
	}

	LaplaceDist := distuv.Laplace{Mu: 0, Scale: 1.0 / eps, Src: rand.NewSource(100)}
	//UniformDist := rand.New(rand.NewSource(101))

	log.Printf("padding num %v", paddingNum)

	for i := uint64(0); i < hashTableSize; i++ {
		fakeQueryNum := int(LaplaceDist.Rand() + float64(paddingNum))
		if fakeQueryNum < 0 {
			fakeQueryNum = 0
		}
		//fakeQueryNum = max(fakeQueryNum, 0)
		for j := 0; j < fakeQueryNum; j++ {
			//fakeQuery := UniformDist.Uint64() % hashTableSize;
			queryBuffer = append(queryBuffer, pb.CuckooBucketQuery{BucketId: i})
		}
	}

	log.Printf("In total : %v packages", len(queryBuffer))

	rand.Seed(uint64(time.Now().UnixNano()))
	rand.Shuffle(len(queryBuffer), func(i, j int) {
		queryBuffer[i], queryBuffer[j] = queryBuffer[j], queryBuffer[i]
	})

	cnt := 0
	for _, q := range queryBuffer {
		//var q = pb.CuckooBucketQuery{BucketId: uint64(i % 10)}
		cnt += 1
		if err := stream.Send(&q); err != nil {
			log.Fatalf("Failed to send a query at %v-th package: %v", cnt, err)
		}
	}

	/*
	for i := 0; i < 10000000; i++ {
		var q = pb.CuckooBucketQuery{BucketId: uint64(i % 10)}
		if err := stream.Send(&q); err != nil {
			log.Fatalf("Failed to send a query at %v-th package: %v", i, err)
		}
	}
	*/
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

func runHashTableInfoQuery(client pb.QueryServiceClient) {
	//ctx, cancel := context.WithTimeout(context.Background(), 1 * time.Second)
	//defer cancel()
	ctx := context.Background()

	in, err := client.GetHashTableInfo(ctx, &pb.HashTableInfoQuery{Dummy: 0})

	if err != nil {
		log.Fatalf("failed to query %v", err)
	}

	hashTableSize = in.Size
	log.Printf("hash table size %v", in.Size)
}


func main() {
	paddingNum = calcLaplacePadding(eps, delta)

	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Failed to connect server %v", address)
	}

	defer conn.Close()
	c := pb.NewQueryServiceClient(conn)

	runHashTableInfoQuery(c);

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