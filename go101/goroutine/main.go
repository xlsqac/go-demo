// package main
package main

import (
	"fmt"
	"runtime"
	"time"
)

func testGoroutineThread() {
	time.Sleep(10 * time.Second)
}

func main() {
	fmt.Println(runtime.NumCPU())
	//defaultGOMAXPROCS := runtime.GOMAXPROCS(0)
	//fmt.Printf("Default GOMAXPROCS: %d\n", defaultGOMAXPROCS)
	//
	for i := 0; i < 4001; i++ {
		go testGoroutineThread()
	}
	//// 设置 GOMAXPROCS 的值
	//runtime.GOMAXPROCS(2)
	//newGOMAXPROCS := runtime.GOMAXPROCS(0)
	//fmt.Printf("New GOMAXPROCS: %d\n", newGOMAXPROCS)
	time.Sleep(30 * time.Second)
}
