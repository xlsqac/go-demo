package main

import "fmt"

type TZ int

func main() {
	var a, b TZ = 1, 2
	c := a + b
	fmt.Printf("%d", c)
}
