package main

import (
	"fmt"
	"sync"
)

type n struct {
	sync.RWMutex
}

func main() {

	m := map[string]int{
		"A": 21,
		"B": 22,
		"C": 23,
	}

	counter := 0
	for k, v := range m {
		if counter == 0 {
			delete(m, "A")
		}
		counter++
		fmt.Println("k, v", k, v)
	}
	fmt.Println(counter)

	m1 := map[string]int{
		"A": 21,
		"B": 22,
		"C": 23,
	}
	counter1 := 0
	for k, v := range m1 {
		if counter1 == 0 {
			m["D"] = 24
		}
		counter1++
		fmt.Println(k, v)
	}

	ch := make(chan int, 10)
	ch <- 1
	ch <- 1
	ch <- 1
	ch <- 1
	for value := range ch {
		fmt.Println(value)
	}

}
