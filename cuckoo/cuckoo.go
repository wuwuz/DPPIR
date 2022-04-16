package cuckoo

import (
	"hash/fnv"
	"log"

	//////////////////"io"
	"encoding/binary"
	"math"
)

const (
	BucketSize        = 4
	secondHashPadding = 0x71516198ffaa77
	emptySlot         = 0
)

type bucket struct {
	Slot [BucketSize]uint64
}

type Cuckoo struct {
	Size    uint64
	Logsize uint64
	Buckets []bucket

	Load uint64
}

func alloc(n uint64) []bucket {
	return make([]bucket, n, n)
}

func NewCuckooHashTable(itemNum uint64) *Cuckoo {
	logsize := uint64(math.Ceil(math.Log2(float64(itemNum / BucketSize * 10 / 9))))
	size := uint64(1) << logsize

	c := &Cuckoo{
		Size:    size,
		Logsize: logsize,
		Buckets: alloc(size),
	}

	return c
}

func NewCuckooHashTableGivenLogSize(logsize uint64) *Cuckoo {
	//logsize := uint64(math.Ceil(math.Log2(float64(itemNum / BucketSize * 10 / 9))))
	size := uint64(1) << logsize

	c := &Cuckoo{
		Size:    size,
		Logsize: logsize,
		Buckets: alloc(size),
	}

	return c
}

func DefaultHash(key uint64) uint64 {
	hash := fnv.New64a()
	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, key)
	hash.Write(b)
	return hash.Sum64()
}

func GetTwoHash(key uint64) (uint64, uint64) {
	return DefaultHash(key), DefaultHash(key ^ secondHashPadding)
}

func (c *Cuckoo) TryInsertToBucket(key uint64, bucketId uint64) bool {
	for i := 0; i < BucketSize; i++ {
		if c.Buckets[bucketId].Slot[i] == 0 {
			c.Buckets[bucketId].Slot[i] = key
			return true
		}
	}
	return false
}

func (c *Cuckoo) TryLookupInBucket(key uint64, bucketId uint64) bool {
	for i := 0; i < BucketSize; i++ {
		if c.Buckets[bucketId].Slot[i] == key {
			return true
		}
	}
	return false
}

func (c *Cuckoo) Insert(key uint64) bool {
	//key = DefaultHash(key)

	if c.Lookup(key) == true {
		return true
	}

	h1, h2 := GetTwoHash(key)
	h1 = h1 % c.Size
	h2 = h2 % c.Size

	if c.TryInsertToBucket(key, h2) {
		c.Load += 1
		return true
	}

	for i := 0; i < 100; i++ {
		if c.TryInsertToBucket(key, h1) {
			c.Load += 1
			return true
		}
		kicked_key := c.Buckets[h1].Slot[i%4]
		c.Buckets[h1].Slot[i%4] = key
		key = kicked_key

		t1, t2 := GetTwoHash(key)
		t1 = t1 % c.Size
		t2 = t2 % c.Size

		if t1 == h1 {
			// the next position is h1, and t1
			h1 = t2
			h2 = t1
		} else {
			h1 = t1
			h2 = t2
		}
	}

	return false
}

func (c *Cuckoo) Lookup(key uint64) bool {
	//key = DefaultHash(key)

	h1, h2 := GetTwoHash(key)
	h1 = h1 % c.Size
	h2 = h2 % c.Size
	if c.TryLookupInBucket(key, h1) || c.TryLookupInBucket(key, h2) {
		return true
	}

	return false
}

func (c *Cuckoo) PrintInfo() {
	log.Printf(
		"Cuckoo hash table: %v buckets, %v slots, %v load factor",
		c.Size,
		BucketSize,
		float64(c.Load)/float64(BucketSize*c.Size)) // / float64(BucketSize * c.size))
}

func (c *Cuckoo) GetBucketCopy(bucketId uint64) bucket {
	var ret bucket
	ret.Slot[0] = c.Buckets[bucketId].Slot[0]
	ret.Slot[1] = c.Buckets[bucketId].Slot[1]
	ret.Slot[2] = c.Buckets[bucketId].Slot[2]
	ret.Slot[3] = c.Buckets[bucketId].Slot[3]
	return ret
}
