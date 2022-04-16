package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	pb "example.com/query"
	"google.golang.org/grpc"

	"golang.org/x/exp/rand"
	"gonum.org/v1/gonum/stat/distuv"

	"example.com/cuckoo"
)

const (
	address = "localhost:50051"
	//address = "10.108.0.3:50051"
	//address      = "localhost:50051"
	//eps          = 0.5
	//delta        = 1e-6
	subBatchSize = 40000
)

var eps float64
var delta float64
var listSize uint64
var realQueryNum uint64
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

func createSingleQuery(bucketId uint64, rng *rand.Rand) *pb.CuckooBucketQuery {
	uniqueId := rng.Uint64()
	//uniqueId := uint64(0)
	//oneTimePad := []uint64{rng.Uint64(), rng.Uint64(), rng.Uint64(), rng.Uint64()}
	oneTimePad_0 := rng.Uint64()
	oneTimePad_1 := rng.Uint64()
	oneTimePad_2 := rng.Uint64()
	oneTimePad_3 := rng.Uint64()
	//oneTimePad_0 := uint64(0)
	//oneTimePad_1 := uint64(0)
	//oneTimePad_2 := uint64(0)
	//oneTimePad_3 := uint64(0)
	/*
		for i := 0; i < cuckoo.BucketSize; i++ {
			oneTimePad[i] = rng.Uint64()
		}
	*/
	return &pb.CuckooBucketQuery{
		BucketId:     bucketId,
		UniqueId:     uniqueId,
		OneTimePad_0: oneTimePad_0,
		OneTimePad_1: oneTimePad_1,
		OneTimePad_2: oneTimePad_2,
		OneTimePad_3: oneTimePad_3,
	}
}

type RealQuery struct {
	Item          uint64
	Result        bool
	BucketQuery_0 pb.CuckooBucketQuery
	BucketQuery_1 pb.CuckooBucketQuery
}

func handleSingleResponse(in *pb.CuckooBucketResponse, queryUniqueIdToRealQuery *map[uint64]*RealQuery) {
	uniqueId := in.UniqueId
	if realQuery, ok := (*queryUniqueIdToRealQuery)[uniqueId]; ok {
		var matchedBucketQuery *pb.CuckooBucketQuery
		if realQuery.BucketQuery_0.UniqueId == uniqueId {
			matchedBucketQuery = &realQuery.BucketQuery_0
		} else {
			//if realQuery.BucketQuery_1.UniqueId == uniqueId {
			matchedBucketQuery = &realQuery.BucketQuery_1
		}
		/*
			if realQuery.Item == 40000 {
				log.Printf("now item = %v ---------------------", realQuery.Item)
			}
		*/

		if matchedBucketQuery.OneTimePad_0^in.Bucket_0 == realQuery.Item ||
			matchedBucketQuery.OneTimePad_1^in.Bucket_1 == realQuery.Item ||
			matchedBucketQuery.OneTimePad_2^in.Bucket_2 == realQuery.Item ||
			matchedBucketQuery.OneTimePad_3^in.Bucket_3 == realQuery.Item {

			if realQuery.Item == 40000 {
				log.Printf("true item = %v -------------------", realQuery.Item)
			}
			realQuery.Result = true
		}
	} else {
		// it's a fake query and its unique id has no registration
		return
	}
}

func runContinuousQuery(client pb.QueryServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(10000*time.Millisecond))
	defer cancel()

	stream, _ := client.ContinuousQuery(ctx)

	startTime := time.Now()

	queryUniqueIdToRealQuery := make(map[uint64]*RealQuery)
	realQueryBuffer := make([]RealQuery, 0, 200000)

	totalReceivedResponse := 0

	//waitc := make(chan struct{})
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		cnt := 0
		for {
			in, err := stream.Recv()
			if err == io.EOF {
				// read done.
				//close(waitc)
				log.Printf("received response cnt %v", cnt)
				//log.Printf("received response cnt %v", cnt)
				log.Printf("real query buffer size %v", len(realQueryBuffer))
				for _, realQuery := range realQueryBuffer {
					if realQuery.Item%uint64(1000) == 0 {
						log.Printf("Item: %v, Result: %v", realQuery.Item, realQuery.Result)
					}
				}

				totalReceivedResponse = cnt
				return
			}
			if err != nil {
				log.Fatalf("Failed to receive a response %v", err)
			}
			cnt += int(in.ResponseNum)
			if cnt%(1000*subBatchSize) == 0 {
				//log.Printf("Got message %s at point(%d, %d)", in.Message, in.Location.Latitude, in.Location.Longitude)
				log.Printf("Got %v-th bucket: ", cnt)
				//for i := 0; i < 1; i++ {
				//	log.Printf("%v ", in.PackedBucket[i])
				//	}
			}

			for i := 0; i < int(in.ResponseNum); i++ {
				handleSingleResponse(in.BatchedResponse[i], &queryUniqueIdToRealQuery)
			}
		}
	}()

	//timer set
	// generate some real queries and put them into a hash table

	go func() {
		defer wg.Done()
		rng := rand.New(rand.NewSource(101))

		queryBuffer := make([]*pb.CuckooBucketQuery, 0, 1000000)

		for i := uint64(1); i <= realQueryNum; i++ {
			//for i := 100000 - 10; i < 100000; i++ {
			item := i
			h0, h1 := cuckoo.GetTwoHash(uint64(item))
			h0 = h0 % hashTableSize
			h1 = h1 % hashTableSize

			realQueryBuffer = append(realQueryBuffer, RealQuery{
				Item:          uint64(item),
				Result:        false,
				BucketQuery_0: *createSingleQuery(h0, rng),
				BucketQuery_1: *createSingleQuery(h1, rng),
			})

			currentQuery := &realQueryBuffer[len(realQueryBuffer)-1] // the current one

			queryUniqueIdToRealQuery[currentQuery.BucketQuery_0.UniqueId] = currentQuery
			queryUniqueIdToRealQuery[currentQuery.BucketQuery_1.UniqueId] = currentQuery

			//queryBuffer = append(queryBuffer, pb.CuckooBucketQuery{BucketId: h1})
			//queryBuffer = append(queryBuffer, pb.CuckooBucketQuery{BucketId: h2})
			queryBuffer = append(queryBuffer, &currentQuery.BucketQuery_0)
			queryBuffer = append(queryBuffer, &currentQuery.BucketQuery_1)
		}

		LaplaceDist := distuv.Laplace{Mu: 0, Scale: 1.0 / eps, Src: rand.NewSource(100)}

		log.Printf("padding num %v", paddingNum)
		for i := uint64(0); i < hashTableSize; i++ {
			//for i := uint64(0); i < uint64(0); i++ {
			fakeQueryNum := int(LaplaceDist.Rand() + float64(paddingNum))
			if fakeQueryNum < 0 {
				fakeQueryNum = 0
			}
			//fakeQueryNum = max(fakeQueryNum, 0)
			for j := 0; j < fakeQueryNum; j++ {
				//fakeQuery := UniformDist.Uint64() % hashTableSize;
				//queryBuffer = append(queryBuffer, pb.CuckooBucketQuery{BucketId: i})
				queryBuffer = append(queryBuffer, createSingleQuery(i, rng))
			}
		}

		log.Printf("In total : %v packages", len(queryBuffer))

		rand.Seed(uint64(time.Now().UnixNano()))
		/*
			rand.Shuffle(len(queryBuffer), func(i, j int) {
				queryBuffer[i], queryBuffer[j] = queryBuffer[j], queryBuffer[i]
			})
		*/

		cnt := 0
		for i := 0; i < len(queryBuffer); i += subBatchSize {
			j := i + subBatchSize
			if j > len(queryBuffer) {
				j = len(queryBuffer)
			}
			/*
				for ; j < i+subBatchSize && j < len(queryBuffer); j++ {
					q = append(q, queryBuffer[j])
				}
			*/
			currentSubBatchSize := j - i
			log.Printf("%v current batch size", currentSubBatchSize)
			cnt += currentSubBatchSize
			if err := stream.Send(&pb.BatchedCuckooBucketQuery{
				QueryNum:     uint64(currentSubBatchSize),
				BatchedQuery: queryBuffer[i:j],
			}); err != nil {
				log.Fatalf("Failed to send a query at %v-th package: %v", cnt, err)
			}

			if cnt%(1000*subBatchSize) == 0 {
				log.Printf("Sent a query at %v-th package", cnt)
			}
		}
		stream.CloseSend()
	}()
	/*
		for _, q := range queryBuffer {
			//var q = pb.CuckooBucketQuery{BucketId: uint64(i % 10)}
			cnt += 1
			if err := stream.Send(&q); err != nil {
				log.Fatalf("Failed to send a query at %v-th package: %v", cnt, err)
			}
		}
	*/

	/*
		for i := 0; i < 10000000; i++ {
			var q = pb.CuckooBucketQuery{BucketId: uint64(i % 10)}
			if err := stream.Send(&q); err != nil {
				log.Fatalf("Failed to send a query at %v-th package: %v", i, err)
			}
		}
	*/
	//<-waitc
	wg.Wait()

	file, err := os.OpenFile("output.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("File does not exists or cannot be created")
		os.Exit(1)
	}
	defer file.Close()

	duration := time.Since(startTime)
	log.Printf("Total time = %v ms", duration.Milliseconds())
	w := bufio.NewWriter(file)

	realQueryNum := len(realQueryBuffer)
	batchTime := float64(duration.Milliseconds()) / 1000
	throughput := float64(realQueryNum) / float64(batchTime)
	overhead := float64(totalReceivedResponse)/float64(realQueryNum*2) - 1

	fmt.Fprintf(w, "eps, delta, listSize, realQueryNum, batchTime(s), throughput(query/s), overheadRatio\n")
	fmt.Fprintf(w, "%.2f, %.2f, %v, %v, %v, %.2f, %.2f\n",
		eps, delta, listSize, realQueryNum, batchTime, throughput, overhead,
	)

	w.Flush()
}

func runSingleQuery(client pb.QueryServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Millisecond*500))
	defer cancel()

	rng := rand.New(rand.NewSource(103))
	for i := 0; i < 1000000; i++ {
		//bucket, err := client.SingleQuery(ctx, &pb.CuckooBucketQuery{QueryNum: 1, BucketId: [uint64(i % 10)]})
		_, err := client.SingleQuery(ctx, createSingleQuery(uint64(i%10), rng))

		if err != nil {
			log.Fatalf("failed to query %v", err)
		}
		if i%1000 == 0 {
			log.Printf("Got %v-th bucket: ", i)
			/*
				for j := 0; j < 1; j++ {
					log.Printf("%v ", bucket.Bucket[j])
				}
			*/
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
	paddingNum = calcLaplacePadding(eps/2.0, delta)

	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Failed to connect server %v", address)
	}

	defer conn.Close()
	c := pb.NewQueryServiceClient(conn)

	runHashTableInfoQuery(c)

	//runSingleQuery(c);
	runContinuousQuery(c)

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
