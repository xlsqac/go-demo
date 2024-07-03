package main

import "fmt"

var p *int

func foo() (*int, error) {
	var i int = 5
	return &i, nil
}

func bar() {
	//use p
	fmt.Println(*p)
}

func main() {
	//{
	//	i := 1
	//	fmt.Println(i)
	//	i := 2
	//}
	{
		fmt.Println(p)
		i := 1
		p := &i
		fmt.Println(p)
	}
	// 修改成这样
	//var err error
	//p, err = foo()
	p, err := foo() // main 中的 p 覆盖了全局变量 p
	if err != nil {
		fmt.Println(err)
		return
	}
	bar() // 对全局变量 p 解引用，全局变量 p 没有被赋值，对于值为 nil 的指针类型解引用会触发 panic
	fmt.Println(*p)
}
