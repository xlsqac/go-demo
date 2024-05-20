package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Name  string
	_name string
}

func (u User) Call() {
	fmt.Println("Call")
}

func (u User) call() {
	fmt.Println("call")
}

func (u User) _call() {
	fmt.Println("_call")

}

func main() {
	u := User{Name: "Name", _name: "_name"}

	v := reflect.ValueOf(u)
	t := reflect.TypeOf(u)

	for i := 0; i < v.NumField(); i++ {
		fmt.Println(v.Field(i))
	}

	for i := 0; i < t.NumMethod(); i++ {
		fmt.Println(t.Method(i))
	}
}
