// package main
// 探究 golang 程序的初始化顺序
package main

import (
	"fmt"
	_ "go101/init/pkg1"
	_ "go101/init/pkg2"
	"time"
)

func test() {
	fmt.Println("test func")
}

func init() {
	fmt.Println("init2 func")
}

func init() {
	fmt.Println("init1 func")
	fmt.Println(Pvar)
}

var Pvar = "pvar"

func main() {
	fmt.Println("main func")
	test()
	//Filea()
	//Eilea()
	time.Sleep(3)
}
