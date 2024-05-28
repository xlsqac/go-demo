package main

import (
	"errors"
	"fmt"
)

func main() {
	var ch chan string
	err := test()
	if err != nil {
		fmt.Println(err)
	}
}

func test() error {
	var err error
	defer func() {
		if r := recover(); r != nil {
			err = errors.New(fmt.Sprintf("%s", r))
		}
	}()
	raisePanic()
	return err
}

func raisePanic() {
	panic("发生了错误")
}
