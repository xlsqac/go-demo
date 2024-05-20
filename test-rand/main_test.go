package main

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

var seed = rand.New(rand.NewSource(time.Now().UnixNano()))

func BenchmarkName(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			fmt.Println(seed)
			seed.Intn(32)
		}
	})
}
