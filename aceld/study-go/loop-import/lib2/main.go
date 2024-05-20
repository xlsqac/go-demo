package lib2

import "fmt"

import "aceld/study-go/loop-import/lib1"

func ImportTest() {
	fmt.Println("package lib2 import test")
}

func init() {
	fmt.Println("package lib2 init")
	lib1.ImportTest()
}
