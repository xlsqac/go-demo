// package main
// 测试 sync.map 和 map 的区别
package main

import (
	"fmt"
	"sync"
)

func main() {
	// 测试两个协程并发读一个 map
	m := make(map[string]string)
	m["a"] = "a"
	m["b"] = "b"
	m["c"] = "c"
	m["d"] = "d"

	var wg sync.WaitGroup

	{
		wg.Add(2)
		go func() {
			for _, v := range []string{"a", "b"} {
				value := m[v]
				fmt.Println(value)
			}
			wg.Done()
		}()

		go func() {
			for _, v := range []string{"c", "d"} {
				value := m[v]
				fmt.Println(value)
			}
			wg.Done()
		}()
	}
	wg.Wait()
}
