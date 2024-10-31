// package main
// https://x.com/func25/status/1846388380386709766
package main

import "fmt"

type Greeter interface {
	Hello()
}

type A struct{}

func (A) Hello() {
	fmt.Println("Hello from A!")
}

func Greet(g any) {
	if v, ok := g.(Greeter); ok {
		v.Hello()
		return
	}
	fmt.Println("g does not implement Greeter!")
}

type NoHello struct{}

func (NoHello) Hello() {
	panic("not implemented")
}

type ANoHello struct {
	NoHello
	A
}

func main() {
	Greet(ANoHello{A: A{}})
}
