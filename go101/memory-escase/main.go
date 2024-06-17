// package main
// 内存逃逸
package main

import (
	"math/rand"
)

func LessThan8192() {
	nums := make([]int, 100) // = 64KB
	for i := 0; i < len(nums); i++ {
		nums[i] = rand.Int()
	}
}

func MoreThan8192() {
	nums := make([]int, 1000000) // = 64KB
	for i := 0; i < len(nums); i++ {
		nums[i] = rand.Int()
	}
}

func NonConstant() {
	s := make([]int, 10)
	for i := 0; i < len(s); i++ {
		s[i] = i
	}
}

func main() {
	NonConstant()
	MoreThan8192()
	LessThan8192()
}
