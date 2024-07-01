// package main
// 单核 CPU 开两个协程，其中一个死循环会怎么样？
// go version: go version go1.22.1 darwin/amd64
// 两个协程，一个是主协程，一个是新开的死循环协程
package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(1) // 将机器设置成单核
	startTime := time.Now()

	// 新开的协程
	go func() {
		//死循环
		for {
		}
	}()
	time.Sleep(time.Millisecond) // 这里加个睡眠，让协程能起来
	elapsed := time.Since(startTime)
	fmt.Printf("程序运行时间: %s\n", elapsed)
	fmt.Println("main goroutine exit")
}
