// package main
package main

import "fmt"

func main() {
	var s [][32]int
	// ok，len() 函数没有 channel 操作，也没有函数调用，不会对 v 求值，在编译阶段就能确定 s[99999]的长度
	fmt.Println(len(s[999999]))
	// error
	fmt.Println(s[999999])
}
