package main

import "fmt"

type SingInterface interface {
	Sing()
}

type TestInterface interface{ TestInterface }

type Chicken struct {
	Name string
}

func (c *Chicken) Sing(i any) {
	fmt.Println(c.Name, "在唱歌")
}

type Cat struct {
	Name string
}

func (c *Cat) Sing() {
	fmt.Println(c.Name, "在唱歌")
}

func sing(s SingInterface) {
	s.Sing()
	c, ok := s.(*Cat)
	fmt.Println(c, ok)
}

func main() {
	cat := &Cat{Name: "cat"}
	chicken := &Chicken{Name: "ikun"}

	sing(cat)
	sing(chicken)

}
