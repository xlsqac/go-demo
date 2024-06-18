// package main
package main

import (
	"fmt"
	"unsafe"
)

type Slice []bool

func (s Slice) Length() int {
	return len(s)
}

func (s Slice) Modify(i int, x bool) {
	s[i] = x // panic if s is nil
}

func (p *Slice) DoNothing() {
}

func (p *Slice) Append(x bool) {
	*p = append(*p, x) // 如果p为空指针，则产生一个恐慌。
}

// typeUnsafeOfNil 各类型 nil 值的尺寸
func typeUnsafeOfNil() {
	var p *struct{} = nil
	fmt.Println(unsafe.Sizeof(p)) // 8

	var s []int = nil
	fmt.Println(unsafe.Sizeof(s)) // 24

	var m map[int]bool = nil
	fmt.Println(unsafe.Sizeof(m)) // 8

	var c chan string = nil
	fmt.Println(unsafe.Sizeof(c)) // 8

	var f func() = nil
	fmt.Println(unsafe.Sizeof(f)) // 8

	var i interface{} = nil
	fmt.Println(unsafe.Sizeof(i)) // 16
}

func main() {
	{
		typeUnsafeOfNil()
	}

	{
		//var _ = (*int)(nil) == (*bool)(nil)
		//var _ = (chan int)(nil) == (chan bool)(nil)
	}

	{
		type Inter *int
		var _ = (Inter)(nil) == (*int)(nil)

		var _ = (interface{})(nil) == (*int)(nil)

		var _ = (chan int)(nil) == (chan<- int)(nil)
		var _ = (chan int)(nil) == (<-chan int)(nil)
	}

	{
		//var _ = ([]int)(nil) == ([]int)(nil)
		//var _ = (map[string]int)(nil) == (map[string]int)(nil)
		//var _ = (func())(nil) == (func())(nil)
	}

	{
		var _ = ([]int)(nil) == nil
		var _ = (map[string]int)(nil) == nil
		var _ = (func())(nil) == nil
	}

	{
		fmt.Println((interface{})(nil) == (*int)(nil)) // false
	}

	{
		fmt.Println((map[string]int)(nil)["key"]) // 0
		fmt.Println((map[int]bool)(nil)[123])     // false
		fmt.Println((map[int]*int64)(nil)[123])   // <nil>
	}

	{
		_ = ((Slice)(nil)).Length
		_ = ((Slice)(nil)).Modify
		_ = ((*Slice)(nil)).DoNothing
		_ = ((*Slice)(nil)).Append

		_ = ((Slice)(nil)).Length()
		((*Slice)(nil)).DoNothing()
	}

	{
		fmt.Println("range nil slice or map")
		var s []int = nil
		var m map[string]int = nil

		for v, ok := range s {
			print(v, ok)
		}

		for v, ok := range m {
			print(v, ok)
		}
	}

	{
		fmt.Println("range nil array")
		var a *[10]int = nil
		// panic: runtime error: invalid memory address or nil pointer dereference
		//for v, ok := range a {
		//	fmt.Println(v, ok)
		//}

		// ok 0~9
		for v := range a {
			fmt.Println(v)
		}
	}

	{
		//fmt.Println("range nil chan")
		//var c chan int = nil
		//for v := range chan bool(nil) {
		//	fmt.Println(v)
		//}
	}

	{
		fmt.Println(*new(*int) == nil)
		fmt.Println(*new([]int) == nil)
		fmt.Println(*new(map[int]bool) == nil)
		fmt.Println(*new(chan string) == nil)
		fmt.Println(*new(func()) == nil)
		fmt.Println(*new(interface{}) == nil)
	}
}
