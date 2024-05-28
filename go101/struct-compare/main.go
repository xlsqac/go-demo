// package main 比较结构体

package main

import (
	"fmt"
)

type T struct {
	Name string
	Age  int
}

func main() {
	t1 := T{
		"name",
		1,
	}

	t2 := T{
		Name: "name"}

	fmt.Println(t1 == t2)
}
