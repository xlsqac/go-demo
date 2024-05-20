package lib1

import "fmt"

import "aceld/study-go/loop-import/lib2"

func ImportTest() {
	fmt.Println("package lib1 import test")
}

func init() {
	fmt.Println("package lib1 init")
	lib2.ImportTest()
}
