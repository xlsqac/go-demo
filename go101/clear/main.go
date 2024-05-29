// package main
// 测试 clear 内置函数用法以及细节
package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4, 5}
	clear(s)
	fmt.Println(s) // [0, 0, 0]

	a := [5]int{1, 2, 3, 4, 5}
	clear(a[1:3])
	fmt.Println(a) // [1, 0, 0, 0, 5]

	m := map[float64]float64{}
	x := 0.0
	m[x] = x
	x /= x
	m[x] = x
	fmt.Println(m)

	for k := range m {
		fmt.Println(k)
		delete(m, k)
	}
	fmt.Println(m)
	clear(m)
	fmt.Println(len(m))
}
