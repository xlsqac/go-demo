package main

import "fmt"

func main() {
	var n int16 = 32767
	var m int32

	var a int32 = 32768
	var b int16
	m = int32(n)

	b = int16(a)

	fmt.Printf("32 bit int is %d\n", m)
	fmt.Printf("16 bit int is %d\n", n)
	fmt.Printf("16 bit int is %d\n", b)
}
