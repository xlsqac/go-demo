package main

import (
	"fmt"
	"reflect"
)

type People struct{}

func (p *People) ShowA() {
	fmt.Println(p)
	fmt.Println("ShowA")
	p.ShowB()
}

func (p *People) ShowB() {
	fmt.Println("ShowB")
}

type Teacher struct {
	People
}

func (t *Teacher) ShowB() {
	fmt.Println("Teacher ShowB")
}

func main() {
	t := &Teacher{}
	fmt.Println(t)
	t.People.ShowA()

	t1 := reflect.TypeOf(t)
	for i := 0; i < t1.NumMethod(); i++ {
		fmt.Println(t1.Method(i))
	}
}
