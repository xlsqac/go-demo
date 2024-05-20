package main

import (
	"math/rand"
	"testing"
	"time"
)

var seed = rand.New(rand.NewSource(time.Now().UnixNano()))

func BenchmarkName(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			seed.Intn(32)
		}
	})
}

func main() {
	BenchmarkName(&testing.B{})
}
