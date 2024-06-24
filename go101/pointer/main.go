// package main
package main

import "fmt"

const A int = 1

func test() int {
	return 1
}

func main() {
	//var p *int
	//// nil 值针不能解引用
	//fmt.Println(*p)
	{
		p0 := new(int)
		fmt.Println(p0)
		fmt.Println(*p0)

		x := *p0
		p1, p2 := &x, &x
		fmt.Println(&x)
		fmt.Println(p1 == p2)
		fmt.Println(p0 == p1)
	}

	{
		type MyInt int64
		type Ta *int64
		type Tb *MyInt

		var pa0 Ta
		var pa1 *int64
		var pb0 Tb
		var pb1 *MyInt

		_ = pa0 == pa1 // 底层类型一致，并且 pa1 是无名类型
		_ = pb0 == pb1 // 同上

		_ = pa0 == nil
		_ = pa1 == nil
		_ = pb0 == nil
		_ = pb1 == nil

		_ = pa0 == pb0     // 都是具名类型
		_ = pa1 == pb1     // 无法隐式转换，都是无名类型
		_ = pa0 == Tb(nil) // 两个都是 nil
	}
}
