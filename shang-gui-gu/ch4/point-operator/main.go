// 其他运算法，指针
package main

import "fmt"

func main() {

	a := 100
	fmt.Println("a 的地址=", &a)

	var ptr *int = &a
	fmt.Println("ptr 指向的值=", *ptr)
}
