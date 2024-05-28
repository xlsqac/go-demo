package main

import (
	"fmt"
	"time"
)

func main() {

	ch := make(chan int)
	// 会死锁
	// go 语句后面的函数调用会先求值再调用，所以 <-ch 是跑在 main goroutine 中的。
	// 发送和接收都在一个 goroutine 中
	//go fmt.Println(<-ch)
	go func() {
		fmt.Println(<-ch)
	}()
	ch <- 5
	time.Sleep(1 * time.Second)

}
