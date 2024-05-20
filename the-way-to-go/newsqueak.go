package main

import (
	"fmt"
	"time"
)

// 向通道输出2开始的自然数序列
func counter(c chan<- int) {
	i := 2
	for {
		c <- i
		i++
	}
}

// 针对listen通道获取的数列，过滤掉是prime倍数的数
// 新的序列输出到send通道
func filter(prime int, listen, send chan int) {
	for {
		if i := <-listen; i%prime > 0 {
			send <- i
		}
	}
}

// 主函数
// 每个通道第一个流出的数必然是素数
// 然后基于这个新的素数构建新的素数过滤器
func sieve() <-chan int {
	c := make(chan int)
	go counter(c)
	prime := make(chan int)
	go func() {
		for {
			p := <-c
			prime <- p
			newc := make(chan int)
			go filter(p, c, newc)
			c = newc
		}
	}()
	return prime
}
func main() {
	prime := sieve()
	for {
		time.Sleep(time.Second)
		select {
		case a := <-prime:
			fmt.Println(a)
		}
	}
}
