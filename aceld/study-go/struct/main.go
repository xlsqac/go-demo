package main

import "fmt"

type Class struct {
	Name string
}

type Student struct {
	Class
	Name string
}

func (s Student) Study() {
	fmt.Printf("%s：正在学习\n", s.Name)
}

func (s Student) SetName(name string) {
	fmt.Printf("methond in %p\n", &s)
	s.Name = name
}

func main() {
	c1 := Class{Name: "终极一班"}
	s1 := &Student{Class: c1, Name: "张三"}
	fmt.Printf("methond in %p\n", s1)
	s1.Study()
	s1.SetName("李四")
	s1.Study()

}
