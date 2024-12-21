package main

import (
	"fmt"
	"github.com/lee20h/bloomfilter-practice/pkg/bloomfilter"
)

func main() {
	bloom := bloomfilter.NewBloomFilter(1000, 5)
	bloom.Add("apple")
	fmt.Println(bloom.Check("apple"))  // true
	fmt.Println(bloom.Check("banana")) // false (허위 긍정 가능성 있음)
}
