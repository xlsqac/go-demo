package main

import "fmt"

func test(b int) (func(), func()) {
	a := func() {
		b++
		fmt.Println(b)
	}
	c := func() {
		fmt.Println(b)
	}
	return a, c
}

func a(b int) func() {
	f1 := func() {
		b++
		fmt.Println(b)
	}
	return f1

}

type N int

func (n N) A() {
	n++
	fmt.Println(n)
}

func (n *N) B() {
	fmt.Println(*n)
}

func main() {
	//var b int = 100
	//a, c := test(b)
	//a()
	//c()
	var d N = 1
	d.A()
	d.B()

	//fmt.Println(b)
}
