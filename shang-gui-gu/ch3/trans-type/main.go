package main

import (
	"fmt"
)

func main() {
	var i1 int64 = 1111111111111111
	var i2 int32 = int32(i1)
	fmt.Println("i1 =", i1)
	fmt.Println("i2 =", i2)

	var n1 int32 = 12
	var n2 int64
	var n3 int8

	// 类型不一样无法相加
	//n2 = n1 + 20
	//n3 = n1 + 20
	// 正确处理
	n2 = int64(n1) + 20
	n3 = int8(n1) + 20
	fmt.Println("n2 =", n2)
	fmt.Println("n3 =", n3)

	// n4 的结果也是越界的但是可以通过编译
	var n4 int8
	n4 = int8(n1) + 127
	fmt.Println("n4 =", n4)

	// n5 越界
	//n5 = int8(n1) + 128
	//var n5 int8
	//fmt.Println("n5 =", n5)
}
