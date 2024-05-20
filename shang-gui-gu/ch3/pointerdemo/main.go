// 指针
// P52
package main

import (
	"Users/xlsqac/Code/go/shang-gui-gu/ch3/str"
	"fmt"
)

func main() {
	var num int = 10
	var ptr *int = &num

	fmt.Printf("ptr的值：%v\n", ptr)
	fmt.Printf("ptr的地址：%v\n", &ptr)
	fmt.Printf("ptr值指向的值：%v\n", *ptr)

	// 通过 ptr 修改 num 的值
	var num1 int = 11
	*ptr = num1
	fmt.Printf("num的值：%v\n", *ptr)
	fmt.Printf("ptr的值：%v\n", ptr)
	fmt.Printf("ptr的地址：%v\n", &ptr)
	fmt.Printf("ptr值指向的值：%v\n", *ptr)

	var num2 int = 1
	var ptr1 *int = &num2
	fmt.Printf("ptr1的值：%v\n", ptr1)

	fmt.Printf("str s:%v", str.s)

}
