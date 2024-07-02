package main

import "fmt"

func main() {
	s := [...]int{1}
	fmt.Println(s)
	a1 := [2]int{}
	a2 := [2]int{}
	fmt.Println(a1 == a2)
	ch1 := make(chan int)
	ch2 := make(chan int)
	fmt.Println(ch1 == ch2) // 输出: false
	//fmt.Println([...]int{1} == [2]int{1})
	//fmt.Println([]int{1} == []int{1})

	a3 := [2][]int{}
	a4 := [2][]int{}
	fmt.Println(a3 == a4)
}
