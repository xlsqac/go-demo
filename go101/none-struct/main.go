// package main
package main

import (
	"fmt"
	"unsafe"
)

var (
	a struct{}
	b struct{}
	c struct{}
	d struct{}
)

func main() {
	println("&a:", &a)
	println("&b:", &b)
	println("&c:", &c)
	println("&d:", &d)

	println("&a == &b:", &a == &b)
	x := &a
	y := &b
	println("x == y:", x == y)

	fmt.Printf("&c(%p) == &d(%p): %t\n", &c, &d, &c == &d)

	type A struct {
		x int
		y string
		z struct{}
	}

	type B struct {
		x int
		z struct{}
		y string
	}

	type C struct {
		z struct{}
		x int
		y string
	}

	a := A{}
	b := B{}
	c := C{}
	fmt.Printf("struct a size: %d\n", unsafe.Sizeof(a))
	fmt.Printf("struct b size: %d\n", unsafe.Sizeof(b))
	fmt.Printf("struct c size: %d\n", unsafe.Sizeof(c))

	var a1 [1000000]string
	var b1 [1000000]struct{}

	fmt.Printf("array a size: %d\n", unsafe.Sizeof(a1))
	fmt.Printf("array b size: %d\n", unsafe.Sizeof(b1))

	var s = make([]string, 1000000)
	var s1 = make([]int, 1000000)
	fmt.Printf("slice s size: %d\n", unsafe.Sizeof(s))
	fmt.Printf("slice s1 size: %d\n", unsafe.Sizeof(s1))
}
