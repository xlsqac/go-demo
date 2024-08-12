// package main
// 简单测试一下 web 服务 p 的数量对性能的提升
package main

import (
	"fmt"
	"log"
	"net/http"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(24)
	http.HandleFunc("/foo", fooHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func fooHandler(w http.ResponseWriter, _ *http.Request) {
	fmt.Println("connect", time.Now())
	time.Sleep(time.Second * 2)
	fmt.Fprintf(w, "Hello, %q", "ok")
}
