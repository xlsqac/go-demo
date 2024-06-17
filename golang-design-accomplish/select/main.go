// package main
// select https://draveness.me/golang/docs/part2-foundation/ch05-keyword/golang-select/
package main

import (
	"fmt"
	"time"
)

func main() {

	// 没有 case 的情况会直接 panic deadlock
	//select {}

	//只有一个 case
	//ch := make(chan int)
	//go func() {
	//	ch <- 0
	//}()
	//select {
	//case v := <-ch:
	//	fmt.Println(v)
	//}

	// 没有可执行的 case 会直接执行 default
	//ch := make(chan int)
	//select {
	//case i := <-ch:
	//	fmt.Println(i)
	//
	//default:
	//	fmt.Println("default")
	//}
	// 编译器会把 select 改成下面的代码
	// selectnbsend 会调用 chansend 并且传非阻塞
	//if selectnbsend(ch, i) {
	//	...
	//} else {
	//	...
	//}
	//
	// 多个 case 且存在 default
	// 总是执行 default
	//ch1 := make(chan int, 1)
	//ch2 := make(chan int)
	//ch1 <- 100

	//go func() {
	//	ch1 <- 0
	//}()
	//go func() {
	//	ch2 <- 0
	//}()
	//select {
	//case v := <-ch1:
	//	fmt.Println(v)
	//case v := <-ch2:
	//	fmt.Println(v)
	//default:
	//	fmt.Println("default")
	//}

	// 随机执行，多个 case 同时就绪时，会随机执行一条
	//select 语句会被 selectgo 改成 if xxx if xxx
	//if 条件的顺序 也称轮询顺序，这个顺序是随机的
	ch := make(chan int)
	go func() {
		for range time.Tick(1 * time.Second) {
			ch <- 0
		}
	}()

	for {
		select {
		case <-ch:
			fmt.Println("case1")
		case <-ch:
			fmt.Println("case2")
		}
	}

	//ch := make(chan int, 2)
	//go func() {
	//	time.Sleep(1 * time.Second)
	//	ch <- 0
	//}()
	//go func() {
	//	time.Sleep(1 * time.Second)
	//	ch <- 1
	//}()

	// 打印的值是随机的，因为写的顺序是不确定的
	//select {
	//case ch <- 1:
	//	fmt.Println("1")
	//case ch <- 0:
	//	fmt.Println("0")
	//}

	// 监听多个 ch，当多个G 都是 0 操作时，总会走最后调用的ch
	// 总是打印 ch1
	//ch1 := make(chan int)
	//ch2 := make(chan int)
	//go func() {
	//	ch2 <- 0
	//}()
	//go func() {
	//	ch1 <- 0
	//}()
	//
	//select {
	//case <-ch1:
	//	fmt.Println("ch1")
	//case <-ch2:
	//	fmt.Println("ch2")
	//}

	// 模拟真实情况下的多个ch
	// 先打印 select，之后的顺序每次执行都是随机的
	//ch1 := make(chan int)
	//ch2 := make(chan int)
	//go func() {
	//	time.Sleep(1 * time.Second)
	//	ch2 <- 0
	//}()
	//go func() {
	//	time.Sleep(1 * time.Second)
	//	ch1 <- 0
	//}()
	//
	//fmt.Println("select")
	//
	//select {
	//case <-ch1:
	//	fmt.Println("ch1")
	//case <-ch2:
	//	fmt.Println("ch2")
	//}

	//ch1 := make(chan int)
	//ch2 := make(chan int)
	//
	//go func() {
	//	for {
	//select {
	//case ch1 <- 1:
	//	fmt.Println("Sent 1 to ch1")
	//case ch2 <- 2:
	//	fmt.Println("Sent 2 to ch2")
	//}
	//	}
	//}()
	//
	//time.Sleep(2 * time.Second)
	//fmt.Println("Received from ch1:", <-ch1)
	//fmt.Println("Received from ch2:", <-ch2)

}
