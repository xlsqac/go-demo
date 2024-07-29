// package main
package main

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"runtime/pprof"
)

func main() {
	fmt.Printf("当前协程数量: %d\n", runtime.NumGoroutine())

	// 创建一个文件来保存堆栈跟踪信息
	f, err := os.Create("goroutine_pprof.txt")
	if err != nil {
		log.Fatal("could not create goroutine profile: ", err)
	}
	defer f.Close()

	// 获取当前所有 goroutine 的堆栈跟踪
	if err := pprof.Lookup("goroutine").WriteTo(f, 1); err != nil {
		log.Fatal("could not write goroutine profile: ", err)
	}

	log.Println("Goroutine profile written to goroutine_pprof.txt")
}
