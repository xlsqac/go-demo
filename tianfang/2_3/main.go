package main

import "fmt"
type People struct {}

func (p *People) showA() {
    fmt.Println("showA")
    p.showB()
}

func (p *People) showB() {
    fmt.Println("showB")
}

type Teacher struct {
    People
}

func (t *Teacher) showB() {
    fmt.Println("teacher showB")
}

func main() {
    t := Teacher{}
    t.showA()
}

