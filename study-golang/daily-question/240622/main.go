package main

import (
	"fmt"
	"time"
)

func main() {
	m := map[int]string{0: "zero", 1: "one"}
	for {
		for k, v := range m {
			fmt.Println(k, v)
		}
		fmt.Println("---------------")
		time.Sleep(time.Second)
	}
}
