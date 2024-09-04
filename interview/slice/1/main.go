// package main
package main

import "fmt"

func modifySlice(s []int) {
	s[0] = 99        // [99 2 3]
	s = append(s, 4) // [99 2 3 4]
	fmt.Println("Inside modifySlice:", s)
}

func main() {
	slice := make([]int, 0, 6) // []
	slice = append(slice, 1)   // [1]
	slice = append(slice, 2)   // [1 2]
	slice = append(slice, 3)   // [1 2 3]
	modifySlice(slice)
	// 1. 为什么在 modiefySlice s[0] = 99 会修改外层的 slice 变量？
	// 传过去的 slice 是一个副本，这个副本和外层的 slice 共享同一个底层数组，s[0] = 99 会修改底层数组的值。 这个值也会反应到外层的 slice 变量中。

	// 2. 既然底层是数组，为什么 append 不会修改外层的 slice 变量？
	// append 函数会返回一个新的切片，这个新的切片是指向底层数组的一个新的切片，这个新的切片的长度是 4，容量是 6。
	// 这个新切片并没有赋值给外层切片
	// append 函数是否会新创建底层数组取决于此次 append 操作是否会触发扩容。

	// 切片的扩容
	// 时机：容量不够时
	// 在 golang 1.18 之后 1024 变成了 256
	// 扩容逻辑：
	// 1. 如果期望容量大于当前容量的两倍，那么直接使用期望容量
	// 2. 如果切片长度小于 1024，直接扩容翻倍
	// 3. 如果切片长度不小于 1024，那么从当前容量开始，每次增加 25% 的容量，直到容量大于期望容量，这里是循环扩容

	// 扩容逻辑的 2 个问题：1. 期望容量怎么来的？2. 3 为什么不直接等于期望容量而是每次扩容 25%。
	// 1.期望容量，期望容量就是当前容量加上 append 的元素数量。
	// 2.不直接扩容到期望容量是要控制扩容的次数。
	fmt.Println("After modifySlice:", slice) // [99 2 3]
}
