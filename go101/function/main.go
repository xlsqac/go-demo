// package main
// go 细节 101，函数声明和调用一章
package main

import (
	"fmt"
)

func funcName(a int) (b int, c int) {
	b = a + 1
	return
}

func funcName1(int, b float64) {
	fmt.Println("1")
}

func funcName2(a int) (int, c int) {
	c = a + 1
	return
}

func funcName3() (r int) {
	var x int
	defer func() {
		r += x
	}()
	x = 10
	return
}

func main() {
	res, b := funcName(1)
	fmt.Println(res, b)
	funcName1(1, 1.0)
	fmt.Println(funcName2(1))
	x := 1
	y := 1

	func() {
		print(x + y)
	}()

	var a = func() {
		print(1)
	}
	a()

	//defer fmt.Println("延迟调用1")
	//defer fmt.Println("延迟调用2")
	//fmt.Println("函数进入退出阶段")
	fmt.Println("\n-------defer--------")

	//defer fmt.Println("9")
	//fmt.Println("0")
	//defer fmt.Println("8")
	//fmt.Println("1")
	//if false {
	//	defer fmt.Println("not reachable")
	//}
	//defer func() {
	//	defer fmt.Println("7")
	//	fmt.Println("3")
	//	defer func() {
	//		fmt.Println("5")
	//		fmt.Println("6")
	//	}()
	//	fmt.Println("4")
	//}()
	//fmt.Println("2")
	fmt.Println(funcName3())

	func() {
		var x = 0
		for i := 0; i < 3; i++ {
			defer func() {
				fmt.Println("b:", i+x)
			}()
		}
		x = 10
	}()

	return
	defer fmt.Println("not reachable")
}

//func main() {
//	var x int
//	defer fmt.Println(x)
//	x = 10
//}
