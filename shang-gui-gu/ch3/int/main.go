package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var num int = 9223372036854775807

	fmt.Println("num =", num)

	var unum uint8 = 255
	fmt.Println("unum =", unum)

	var num64 int64 = 10
	fmt.Printf("num64 bytes %d \n", unsafe.Sizeof(num64))
}
