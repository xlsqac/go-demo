// package main
// 类型内嵌
// 内嵌字段和普通字段不同，内嵌字段字段名和字段类型相同，字段名是隐式的
// 一个类型需要满足以下条件才可以被内嵌
// 类型名 T 既不表示一个具名指针类型也不表示一个基类型为指针类型或接口类型的指针类型
// 一个指针类型 *T 只有在 T 是一个类型名并且 T 既不表示一个指针类型也不表示一个接口类型
// 无名非指针类型不可内嵌
// 一般只有内嵌有字段和方法的类型才有意义
package main

type IntPtr *int

type IntPtrAlias = *int

type IntPtrPtrAlias = *IntPtr

type AliasPtr = *struct {
	name string
	age  int
}

type P = *bool
type M = map[int]int

type AliasPtrInterface = *interface{}

var x int

// S 内嵌非指针类型、接口类型和指针类型
type S struct {
	string      // 具名非指针类型 ok
	error       // 具名接口类型 ok
	*int        // 无名指针类型，T int 是类型名 ok
	P           // 无名指针类型的别名，T bool 是类型名 ok
	M           // 无名非指针类型的别名
	IntPtrAlias // 无名指针类型的别名，T int 是类型名
	AliasPtr    // 无名指针类型的别名，T 是结构体 ok
	//IntPtr            // 具名指针类型 error
	//IntPtrPtrAlias    // 无名指针类型的别名，T 是指针类型 error
	//AliasPtrInterface // 无名指针类型的别名，T 是接口类型 error
	//x // 无名类型 error
	// struct{} // 无名类型 error
	//int // 类型不能重复，*int 的基类型是 int，在这里重复了 error
	//S // 不能内嵌自己，会在定义时 error
}

func main() {
	var s S
	// string 是内嵌的，隐式字段名就是类型名
	s.string = "ok"
}
