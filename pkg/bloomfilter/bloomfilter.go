package bloomfilter

import (
	"hash/fnv"
)

type BloomFilter struct {
	bitArray  []uint64
	size      uint64
	hashCount int
}

// NewBloomFilter 새로운 Bloom filter를 생성합니다.
func NewBloomFilter(size uint64, hashCount int) *BloomFilter {
	return &BloomFilter{
		bitArray:  make([]uint64, (size+63)/64), // 비트 배열 초기화
		size:      size,
		hashCount: hashCount,
	}
}

// hash 주어진 아이템에 대해 해시 값을 생성합니다.
func (b *BloomFilter) hash(item string, seed int) uint64 {
	h := fnv.New64a()
	h.Write([]byte(item))
	h.Write([]byte{byte(seed)}) // 시드를 해시 입력에 추가
	return h.Sum64() % b.size
}

// Add Bloom filter에 아이템을 추가합니다.
func (b *BloomFilter) Add(item string) {
	for i := 0; i < b.hashCount; i++ {
		index := b.hash(item, i)
		bitIndex := index / 64
		bitPosition := index % 64
		b.bitArray[bitIndex] |= (1 << bitPosition) // 해당 비트를 1로 설정
	}
}

// Check 아이템이 Bloom filter에 포함되어 있는지 확인합니다.
func (b *BloomFilter) Check(item string) bool {
	for i := 0; i < b.hashCount; i++ {
		index := b.hash(item, i)
		bitIndex := index / 64
		bitPosition := index % 64
		if (b.bitArray[bitIndex] & (1 << bitPosition)) == 0 {
			return false // 하나라도 0이면 포함되지 않음
		}
	}
	return true // 모든 비트가 1이면 포함됨
}
