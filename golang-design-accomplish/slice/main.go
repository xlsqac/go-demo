// package main
// https://draveness.me/golang/docs/part2-foundation/ch03-datastructure/golang-array-and-slice/
package main

import "fmt"

func NewSlice() []int {
	arr := [3]int{1, 2, 3}
	slice := arr[0:1]
	slice[0] = 33
	fmt.Println(arr)
	return slice
}

func main() {
	s := NewSlice()
	fmt.Println(s)

	s1 := make([]int, 3, 4)
	fmt.Println(cap(s1))
	s1 = append(s1, 1, 2)
	fmt.Println(cap(s1))
	s11 := []int{2, 3, 4, 5, 6}
	s12 := []int{1, 1, 1, 1, 1, 1, 1}
	copy(s12, s11)
	fmt.Println(s12)

}
