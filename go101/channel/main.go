// package main
package main

import (
	"fmt"
	"time"
)

func main() {
	//var c chan int
	//var c1 chan<- int
	//var c2 <-chan int
	//
	//fmt.Println(c == c1)
	//fmt.Println(c1 == c2)
	//c3 := (c2)(c1)

	//{
	//	c1 := make(chan int)
	//	c2 = c1
	//	fmt.Println(c1 == c2)
	//}

	{
		//c3 := make(chan int, 1)
		//fmt.Println(len(c3)) // 0
		//fmt.Println(cap(c3)) // 1
		//c3 <- 1
		//fmt.Println(len(c3)) // 0
		//close(c3)
		//v, sentBeforeClosed := <-c3
		//fmt.Println(v, sentBeforeClosed)
	}
	{
		var ball = make(chan string)
		kickBall := func(playerName string) {
			for {
				fmt.Print(<-ball, "传球", "\n")
				time.Sleep(time.Second)
				ball <- playerName

			}
		}
		go kickBall("张三")
		go kickBall("李四")
		go kickBall("王")
		go kickBall("刘")
		ball <- "裁判"    // 开球
		var c chan bool // 一个零值nil通道
		<-c             // 永久阻塞在此
	}
}
