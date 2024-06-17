// package main
// type alias 类型别名
// type identifier = identifier
// type intAlias = int  // intAlias 是 int 的类型别名
// type alias 的作用：
// 1. 别名可以应用于任意类型，包括其他包中的类型，可以给其他包的类型起一个更符合当前场景的名称提高代码可读性
// type F = func()  ok
// type I = interface() ok
package main

import (
	"fmt"
)

// S 循环定义是不允许的
//type S = struct {
//	name *S1
//}
//
//type S1 = S

// S 别名共享同样的方法集
type S struct{}

func (s S) func1() {

}

type S1 = S

func (s1 S1) func2() {}

// MyReader 不能给别的包的类型通过别名来定义方法
//type MyReader = bufio.Reader
//
//func (m MyReader) AliasMethod() {
//	fmt.Println("This is alias method")
//}

// originalType 可以通过别名灵活的控制导出性
// originalType 是未导出的
// type originalType struct {
//    Field1 string
//    Field2 int
//}

// OriginalType 是导出的
// type OriginalType = originalType

func main() {
	// 字面量是无类型的，无类型可以赋值给类似类型的变量或常量，所以此行代码是合法的
	var x int32 = 32.0
	//var y int = x  // 类型编译不通过，int32 和 int 的内存占用和内存布局就不一样
	//var z rune = x // rune 是 int32 的类型别名，可以编译
	//fmt.Println(y, z)
	fmt.Println(x)

	s1 := S1{}
	s1.func1() // ok
	s1.func1() // ok

	s := S{}
	s.func2() // ok

}
