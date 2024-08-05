// package main
// https://www.cnblogs.com/CodeWithTxT/p/11370215.html
// 测试执行顺序
package main

import "fmt"

func main() {
	done := make(chan bool)
	value := []string{"a", "b", "c"}
	for _, v := range value {
		fmt.Println("---->", v)
		go func(u string) {
			fmt.Println(v)
			done <- true
		}(v)
	}

	for range value {
		<-done
	}
}
