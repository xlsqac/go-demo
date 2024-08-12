package main

import "fmt"

type Student struct {
	Name string
}

var list map[string]Student

func main() {
	list = make(map[string]Student)
	list["student"] = Student{"Aceld"}

	// list["student"] 存的是结构体值的一个拷贝，值引用是不可以修改的
	list["student"].Name = "LDB"
	fmt.Println(list["student"])
}
