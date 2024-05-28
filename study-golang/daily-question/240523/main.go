package main

import "fmt"

type People struct{}

func (p *People) ShowA() {
	fmt.Println("showA")
	p.ShowC()
}
func (p *People) ShowB() {
	fmt.Println("showB")
}

type Teacher struct {
	People
}

func (t *Teacher) ShowB() {
	fmt.Println("teacher showB")
}

func (t *Teacher) ShowC() {
	fmt.Println("teacher showC")
}

func main() {
	t := Teacher{}
	t.ShowA()
}
