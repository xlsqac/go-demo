// package main
// 切片和数组字面量初始化时可以带键
// 键会作为该元素的索引
// 不带键的元素使用之前的元素索引加一
// 没有的索引用空值填充
package main

import "fmt"

func main() {
	// [0, 0, 1]
	// 1 的索引为 2，前面的用 0 填充
	a := []int{2: 1}
	fmt.Println(a)
	fmt.Println(len(a))

	// [0, 77, 88, 0, 44, 55, 66]
	// 44 的索引是 4，55 和 66 就是 5 和 6
	// 77 给了键 1，88就是 2
	// 0 和 3 就缺失了用 0 补充
	var x = []int{4: 44, 55, 66, 1: 77, 88}
	fmt.Println(x)
	fmt.Println(len(x), x[2])

	var intSet = [...]int{2: 2, 0: 4, 6}
	fmt.Println(intSet)

	// 索引重复了，编译不通过
	//
	//var y = []int{4: 44, 55, 66, 3: 77, 88}
	fmt.Printf("%s111\n", "1")

}
