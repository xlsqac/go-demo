package main

import "fmt"

type Integer int

func (a Integer) Add(b Integer) Integer {
	return a + b
}

func main() {
	var a Integer = 1
	var b Integer = 2
	var i interface{} = &a
	value, ok := i.(*Integer)
	fmt.Printf("%d\n", value)
	fmt.Printf("%t\n", ok)
	sum := i.(*Integer).Add(b)
	fmt.Print(sum)
}
