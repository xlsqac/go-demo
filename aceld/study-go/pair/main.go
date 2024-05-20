package main

import "fmt"

func main() {

	var a string
	a = "aceld"

	var allType interface{}
	allType = a

	str, ok := allType.(int)
	fmt.Println(str)
	fmt.Println(ok)

}
