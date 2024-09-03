// package main
// 三个 goroutine，2 个计算（不打印），一个加到 1 万，1 个加到 2 万；第三个等它们两个计算完，打印它们的和。
// https://mp.weixin.qq.com/s/W-BOtDkKO-JaeQKqD_wU1A 问题 5
package main

import (
	"fmt"
	"sync"
)

// main
// add1 和 add2 做计算，计算完成后把结果发送到 chan 中，print 函数从 chan 中读取结果并打印。
// 并且利用 sync.WaitGroup 来保证主 goroutine 等待 print 执行完成后再退出。
func main() {
	var wg = &sync.WaitGroup{}
	sumchan1 := make(chan int)
	sumchan2 := make(chan int)

	go add1(sumchan1)
	go add2(sumchan2)

	wg.Add(1)
	go print(sumchan1, sumchan2, wg)

	wg.Wait()
}

// add1
func add1(sumchan chan int) {
	sum := 0
	for i := 0; i < 10000; i++ {
		sum += 1
	}
	sumchan <- sum
}

// add2
func add2(sumchan chan int) {
	sum := 0
	for i := 0; i < 20000; i++ {
		sum += 1
	}
	sumchan <- sum
}

// print
func print(sumchan1, sumchan2 chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	sum1 := <-sumchan1
	sum2 := <-sumchan2
	fmt.Println(sum1 + sum2)
}
