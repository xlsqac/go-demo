// package main
// 验证 time.After 导致的内存泄露问题
package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"time"
)

func main() {
	//timer := time.NewTimer(time.Second * 5)
	timer1 := time.NewTicker(time.Second * 3)
	defer timer1.Stop()

	//ch := make(chan int)
	//go func() {
	//	in := 1
	//	for {
	//		in++
	//		ch <- in
	//	}
	//}()

	go func() {
		for {
			//timer.Reset(time.Second * 5)
			select {
			//case _ = <-ch:
			//	continue
			case <-timer1.C: // 只会触发一次 没有 reset 的话
				fmt.Println("time out", time.Now().Unix())
			}
		}
	}()
	_ = http.ListenAndServe("0.0.0.0:6060", nil)

}
