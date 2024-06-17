package main

import "fmt"

func f() (int, int) {
	return 1, 1
}

func main() {
	var x int
	//x, _ := f() // error := 左边必须有一个新变量，_ 算
	//x, _ = f()  // ok
	//x, y := f()  // ok y 是新变量
	//x, y = f() // error y 没有声明，需要先声明

	fmt.Println(x)

}
