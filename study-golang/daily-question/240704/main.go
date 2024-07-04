package main

import "fmt"

func main() {
	v := []int{1, 2, 3}
	for i, value := range v {
		fmt.Println(i, value)
		v = append(v, i)
	}
	fmt.Println(v)
}
