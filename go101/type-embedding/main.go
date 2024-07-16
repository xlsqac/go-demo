// package main
// 类型内嵌
// 内嵌字段和普通字段不同，内嵌字段字段名和字段类型相同，字段名是隐式的
// 一个类型需要满足以下条件才可以被内嵌
// 类型名 T 既不表示一个具名指针类型也不表示一个基类型为指针类型或接口类型的指针类型
// 一个指针类型 *T 只有在 T 是一个类型名并且 T 既不表示一个指针类型也不表示一个接口类型
// 无名非指针类型不可内嵌
// 一般只有内嵌有字段和方法的类型才有意义
package main

import (
	"fmt"
)

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

type A struct {
	x string
	y int
}
type B struct {
	x string
}

type C struct {
	A
	B
}

type D struct {
	C
}

type Fooer interface {
	Foo() string
}
type Container struct {
	Fooer
}

type T1 struct {
	test.T
}
type HelloWorld interface {
	hello()
}

func (cont Container) Foo() string {
	fmt.Println("cont Foo")
	return cont.Fooer.Foo()
}

// sink takes a value implementing the Fooer interface.
func sink(f Fooer) {
	fmt.Println("sink:", f.Foo())
}

// TheRealFoo is a type that implements the Fooer interface.
type TheRealFoo struct {
}

func (trf TheRealFoo) Foo() string {
	return "TheRealFoo Foo"
}

type I interface {
	M()
	int
}

type I1 struct {
	s string
}

func (i *I1) M() {
	fmt.Println("M")
}

type i interface {
	int
	error
}
type HelloWorld1 interface {
	hello()
	*hello //ok
	int    // error 没有 int 类型实现这个接口，但是又增加了 int 类型的约束，所以基于此的泛型函数是调用不了的
}

type Person struct {
	name string
	HelloWorld
}

func Print[T HelloWorld1](s []T) {
	for _, v := range s {
		fmt.Println(v)
	}
}

type hello struct {
}

func (h *hello) hello() {
	fmt.Println("hello")
}

func main() {
	var s S
	// string 是内嵌的，隐式字段名就是类型名
	s.string = "ok"
	//s.test = "ok"

	c := C{A: A{x: "a", y: 1}, B: B{x: "b"}}
	d := D{C: c}
	//fmt.Println(c.x)
	fmt.Println(c.A.x)
	fmt.Println(c.B.x)
	fmt.Println(d.y)

	//t1 := T1{}
	//fmt.Println(t1.x)
	co := Container{Fooer: TheRealFoo{}}
	sink(co)

	i := &I1{}
	i.M()
	t1 := T1{}
	fmt.Println(t1.x)
	//t1 := T1{}
	//fmt.Println(t1.x)
	h := &hello{}
	p := &Person{"claude", h}
	p.hello()
	p.HelloWorld.hello()

	slice := []*hello{h}
	Print(slice)
}
