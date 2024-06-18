package main

import (
	"fmt"
)

func main() {
	var x string = nil // error 字符串不能赋值为 nil
	// error
	// 字符串不能和 nil 比较
	if x == nil {
		x = "default"
	}
	fmt.Println(x)
}
