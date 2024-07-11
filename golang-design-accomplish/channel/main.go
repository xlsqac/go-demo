// package main
// 有 4 个 g，编号是 1，2，3，4，每秒会打印自己的编号，使用 channel 使每次打印的顺序都是 1，2，3，4
package main

import (
	"fmt"
	"time"
)

func printNumber(id int, ch chan int, nextCh chan int) {
	for {
		<-ch
		fmt.Println(id)
		time.Sleep(time.Second * 1)
		nextCh <- 1
	}
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)
	ch4 := make(chan int)
	// 启动 4 个 goroutine，每个 goroutine 都从自己的 channel 接收数据，然后打印自己的编号，然后发送数据到下一个 channel

	go printNumber(1, ch1, ch2)
	go printNumber(2, ch2, ch3)
	go printNumber(3, ch3, ch4)
	go printNumber(4, ch4, ch1)

	ch1 <- 1

	select {}

}
