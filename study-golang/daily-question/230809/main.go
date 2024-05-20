package main

import "fmt"

func change(s ...int) []int {
	fmt.Printf("len: %d  cap: %d pointer: %p\n", len(s), cap(s), s)
	s = append(s, 3)
	fmt.Printf("len: %d  cap: %d pointer: %p\n", len(s), cap(s), s)
	return s
}

func main() {
	slice := make([]int, 5)

	slice[0] = 1
	slice[1] = 2
	slice1 := change(slice...)
	fmt.Println(slice)
	fmt.Println(slice1)
	fmt.Println()

	slice2 := change(slice[0:2]...)
	fmt.Println(slice)
	fmt.Println(slice2)
}
