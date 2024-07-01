package main

import "fmt"

type People interface {
	Speak(string) string
}

type Student struct{}

func (s *Student) Speak(think string) (talk string) {
	if think == "speak" {
		talk = "speak"
	} else {
		talk = "hi"
	}
	return
}

func main() {
	//var peo People = Student{} // error 实现 Speak 方法绑定的是指针
	var peo People = &Student{}
	think := "speak"
	fmt.Println(peo.Speak(think))
}
