package main

import "fmt"

type testStruct struct {
	field string
}

func value(t testStruct) {
	t.field = "field1"
}

func pointer(t *testStruct) {
	t.field = "field1"
}

func main() {
	t := &testStruct{field: "testStruct"}
	value(*t)
	fmt.Println("value modify t.field", t.field)
	pointer(t)
	fmt.Println("pointer modify t.field", t.field)
}
