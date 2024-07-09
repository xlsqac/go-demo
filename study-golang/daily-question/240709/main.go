package main

import "fmt"

func main() {
	var a = []int{1, 2, 3, 4, 5}
	var r [5]int

	for i, v := range a {
		if i == 0 {
			a[1] = 12
			a[2] = 13
		}
		r[i] = v
	}
	// a is slice，a [1 12 13 4 5] r [1 12 13 4 5]
	// 切片的数据结构维护了对底层数组的指针，range a 的副本依然指向原来的底层数组，通过 v 可以找到修改后的值
	// a is array a [1 12 13 4 5] r [1 2 3 4 5]
	// 在 range 时会拷贝原数据，修改 a 时，拷贝的数据不会同步修改，所以 r 是 a 修改前的值
	fmt.Println("a", a)
	fmt.Println("r", r)

}
