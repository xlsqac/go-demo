package main

import "fmt"

var (
	a = 1
	b = 2
)

func main() {
	fmt.Println("a=", a)
	fmt.Println("&a=", &a)
	fmt.Println("b=", b)
	a = 3
	testModifyGlobalVar()
}

func testModifyGlobalVar() {
	fmt.Println("&a=", &a)
	fmt.Println("a=", a)
}
