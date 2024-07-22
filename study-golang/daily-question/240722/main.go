// package main
// 在一个协程中遍历 map 的同时删除 k, 可以正常删除, 并且会删除原值
package main

import (
	"fmt"
	"strconv"
	"unsafe"
)

func main() {
	m1 := make(map[string]int)
	fmt.Println(unsafe.Sizeof(m1))

	m := make(map[string]int)

	// 向 map 中添加 1000 个元素
	for i := 0; i < 1000; i++ {
		key := "key" + strconv.Itoa(i)
		m[key] = i
	}

	//fmt.Println(m)
	fmt.Println(unsafe.Sizeof(m))
	for k := range m {
		//fmt.Println(k)
		delete(m, k)
	}
	//fmt.Println(m)
	fmt.Println(unsafe.Sizeof(m))
}
