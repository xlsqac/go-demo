// package main
package main

import "fmt"

// main
// a 和 b 的区别
// a 长度为 0，容量为10 元素类型是 bool，打印为 []
// b 长度为 10，容量为 10，元素类型是 bool，打印为 [false false...] 一共 10 个 false
// bool 的零值是 false
func main() {
	a := make([]bool, 0, 10)
	b := make([]bool, 10, 10)

	fmt.Println(a)
	fmt.Println(b)
}
