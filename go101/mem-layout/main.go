// package main
// 内存布局
package main

import (
	"fmt"
	"unsafe"
)

type T struct {
	// 内存对齐保证是 8n 个字节
	b int64 // 8
	a int8  // 1 // 内存对齐保证 +1
	c int16 // 2 // 为了让 T 的尺寸是内存保证的倍数需要 +4

	// 8+1+1+2+4 =16
}

type T1 struct {
	// 内存对齐保证是 8n 个字节
	a int8  // 1 // 内存对齐保证 +7
	b int64 // 8
	c int16 // 2 // 为了让 T 的尺寸是内存保证的倍数需要 +6

	// 1+7+8+2+6 = 24
	_ [0]func()
}

func main() {
	t := T{}
	fmt.Println(unsafe.Sizeof(t)) // 16 不是 11，因为有内存对齐保证

	t1 := T1{}
	fmt.Println(unsafe.Sizeof(t1)) // 24 不是 11，因为有内存对齐保证
}
