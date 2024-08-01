package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	n := 150
	wg.Add(n)
	t := time.Now()
	for i := 0; i < n; i++ {
		go Work(&wg)
	}
	wg.Wait()
	fmt.Println("end", time.Since(t))
}

func Work(wg *sync.WaitGroup) {
	_, err := http.Get("http://127.0.0.1:8080/foo")
	if err != nil {
		fmt.Println(err)
	}
	wg.Done()
}
