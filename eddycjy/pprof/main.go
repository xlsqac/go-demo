// package main
// https://golang2.eddycjy.com/posts/ch6/01-pprof-1/
// pprof 的使用
package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
	"time"
)

var datas []string

func Add(str string) int {
	data := []byte(str)
	datas = append(datas, string(data))
	return len(datas)
}

func main() {
	go func() {
		for {
			log.Printf("len: %d", Add("go-programming-tour-book"))
			time.Sleep(time.Millisecond * 10)
		}
	}()

	_ = http.ListenAndServe("0.0.0.0:6060", nil)
}
