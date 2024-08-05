// package main
// https://www.cnblogs.com/CodeWithTxT/p/11370215.html
// 查看调度器信息
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("main start")
	var wg sync.WaitGroup
	wg.Add(10)
	fmt.Println("create g start")
	for i := 0; i < 10; i++ {
		go work(&wg)
	}
	fmt.Println("create g end")
	wg.Wait()
	time.Sleep(time.Second * 3)
}

func work(wg *sync.WaitGroup) {
	fmt.Println("g start sleep")
	time.Sleep(time.Second * 1)
	fmt.Println("g sleep end")
	var count int
	for i := 0; i < 1e10; i++ {
		count++
	}
	wg.Done()
	fmt.Println("g end")
}
