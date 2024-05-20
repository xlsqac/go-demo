package main

import (
	"fmt"
	"reflect"
)

func main() {
	// 没有动态类型也没有动态值，所以是 nil
	var i interface{}
	fmt.Println("i", i)
	if i == nil {
		fmt.Println("i nil")
	} else {
		fmt.Println("i not nil")
	}

	// 空接口有了动态类型，所以不为 nil 了
	var b *int = nil
	var i1 interface{} = b
	fmt.Println("i1 type", reflect.TypeOf(i1))
	fmt.Printf("i1 %v\n", i1)
	if i1 == nil {
		fmt.Println("i1 nil")
	} else {
		fmt.Println("i1 not nil")
	}

	var i2 interface{} = b
	fmt.Println("i2 type", reflect.TypeOf(i2))
	if i2 == (*int)(nil) {
		fmt.Println("i2 nil")
		return
	} else {
		fmt.Println("i2 not nil")
	}

}
