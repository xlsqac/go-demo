// package main
// 泛型
package main

import "fmt"

type MyInt int

type num interface {
	~int | float64
}

func add[T num](a, b T) {
	res := a + b
	fmt.Println(res)
}

type S[T any] struct {
	Field1 T
	Field2 string
}

type S1[T any, T1 int] struct {
	Field1 T
	Field2 T1
}

func (s *S1[T, T1]) push() {}

func main() {

	a := 1
	b := 2
	add(a, b)

	c := 1.1
	d := 2.2
	add(c, d)

	//e := "1"
	//f := "2"
	//add(e, f)

	var e MyInt = 1
	var f MyInt = 2
	add(e, f)

	S[int]{1, "1"}
	S[string]{"1", "1"}
	S1[string, int]{"1", 2}
	S1[int, int]{1, 2}
}
