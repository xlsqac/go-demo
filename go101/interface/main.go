// package main
// interface
// https://github.com/studygolang/GCTT/blob/master/published/tech/20180326-Go-Data-Structures-Interface.md
package main

import (
	"fmt"
	"strconv"
)

type Stringer interface {
	String() string
}

type Binary uint64

func (i Binary) String() string {
	return strconv.FormatUint(i.Get(), 2)
}

func (i Binary) Get() uint64 {
	return uint64(i)
}

// ReadWriterCloser
// 基本接口类型，可以做值类型
type ReadWriterCloser = interface {
	Read([]byte) (n int, err error) // 方法描述
	Write([]byte) (n int, err error)
	error                      // 类型名称 基本接口类型
	interface{ Close() error } // 类型字面量形式
}

type read struct{}

func (r read) Read(b []byte) (n int, err error) {}

func (r read) Write(b []byte) (n int, err error) {}

func (r read) Close() error {}

func (r read) Error() string {}

// AnyByteSlice
// 内嵌了一个近似类型
// 他的类型集由所有底层类型为 []byte的类型组成
// 约束接口类型
type AnyByteSlice = interface {
	~[]byte
}

// Unsigned
// 内嵌了一个类型并集, 它的类型集包括 6 个类型
// 约束接口类型
type Unsigned interface {
	uint | uint8 | uint16 | uint32 | uint64 | uintptr
}

// 空接口, 也是基本接口类型
// interface {}

type Nothing interface{}

// 无名接口类型有相同的方法集，那么认为他们是相同类型
type interface1 = interface {
	func1()
}

type interface2 = interface {
	func1()
}

// interface3 的类型集 (S1) 同时是 interface4 的类型集
// interface3 的方法集也是 interface4 的子集
// 可以说 interface3 实现了 interface 4
type interface3 interface {
	func3()
}

type interface4 interface {
	func3()
	func4()
}

type S1 struct{}

func (s S1) func3() {}
func (s S1) func4() {}

func main() {
	b := Binary(200)
	s := Stringer(b)
	fmt.Println(s.String())

	var any1 interface{}
	s1 := any1.(Stringer)
	for i := 0; i < 100; i++ {
		fmt.Println(s1.String())
	}

	// 值包裹
	// 类型S1 实现了 interface4，所以可以隐式转换
	// 将 s2 的一个复制将被包裹在 i 中。操作复杂度为 O(n)，n 是 s2 的尺寸
	var s2 = S1{}
	// s2 是 i 的动态值，s2 的类型 S1 是 i 的动态类型
	// 运行时会分析和构建 i 和 s2 的类型的实现关系信息，并将此信息存到 i 中，最多构建一次
	// 并且此信息会被缓存到内存中的一个全局映射中。
	// 一个非零接口值在内部只是使用一个指针字段来引用全局映射中的一个实现关系信息条目
	// 接口的零值是 nil。
	// 实现关系信息的内容：动态类型的信息，一个方法表（切片类型），存储了接口类型指定的并且此动态类型声明方法。
	// 1. 实现反射的关键。2 是实现多态的关键。
	var i interface4 = s2
	_, ok := i.(S1)
	if ok {
		fmt.Println(ok)
	}
	i = nil
	var i1 interface{} = s2
	var i2 interface3 = s2
	// in3 实现了 in4，O(1), 编译器做过优化
	var i3 interface3 = i

	// 非基本接口类型无法用作接口值
	var u1 Unsigned = int8(1)
	var abs AnyByteSlice = []byte{1}

	// 多态，一个接口值通过包裹不同动态类型的动态值来表现出不同行为
	// 调用一个接口值的方法时实际上将调用此接口值的动态值的方法
	// i.func3() == s2.func3()
	i.func3()
	s2.func3()
	// fmt.Println()
	// 如果没有多态，打印 int 和 string 就得实现两个函数
	// PrintlnInt 和 PrintlnString
	fmt.Println()

	// 反射
}
