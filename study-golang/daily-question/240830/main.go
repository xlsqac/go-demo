package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	st := time.Now()
	var wg sync.WaitGroup
	var total int64
	// var mu sync.Mutex
	wg.Add(10)
	sum := 0
	for i := 1; i <= 10; i++ {
		sum += i
		go func() {
			atomic.AddInt64(&total, int64(i))
			wg.Done()
		}()
	}
	fmt.Println(time.Since(st))
	wg.Wait()
	fmt.Printf("total:%d sum %d\n", total, sum)
	fmt.Println(time.Since(st))
	chanSwith()
}

func chanSwith() {
	ch := make(chan int)
	st := time.Now()
	total, sum := 0, 0
	for i := 1; i <= 10; i++ {
		sum += i
		go func() {
			total += i
			ch <- i
		}()
		<-ch
	}
	fmt.Println(time.Since(st))
	fmt.Printf("total:%d sum %d\n", total, sum)
	fmt.Println(time.Since(st))
}
