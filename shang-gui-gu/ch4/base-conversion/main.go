package main

import "fmt"

func main() {

	var i int = 5
	// 二进制输出
	fmt.Printf("%b\n", i)

	// 八进制，满 8 进 1，以数字 0 开头表示
	var j int = 011
	fmt.Println("j=", j)

	// 十六进制
	var k int = 0x11
	fmt.Println("k=", k)

}
