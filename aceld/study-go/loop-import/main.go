package main

import "fmt"

func main() {

	s := []int{1}
	s1 := s
	fmt.Printf("s addr %p\n", &s)
	fmt.Printf("s1 addr: %p\n", &s1)
	s1[0] = 2
	fmt.Printf("s 0 %d\n", s[0])
	fmt.Printf("s1 0 %d\n", s1[0])
	s2 := append(s1, 3)
	fmt.Printf("s  %d\n", s)
	fmt.Printf("s1 %d\n", s1)
	fmt.Printf("s2 %d\n", s2)
}
