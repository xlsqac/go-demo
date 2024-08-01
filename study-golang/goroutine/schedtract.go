package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("main start")
	//runtime.GOMAXPROCS(5)
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go work(&wg, i)
	}
	wg.Wait()
	// Wait to see the global run queue deplete.
	time.Sleep(3 * time.Second)
}
func work(wg *sync.WaitGroup, i int) {
	fmt.Printf("g start %d\n", i)
	var counter int
	for i := 0; i < 1e10; i++ {
		counter++
	}
	wg.Done()
	fmt.Printf("g end %d\n", i)
}
