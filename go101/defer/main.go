// package main
package main

import "fmt"

func main() {
	defer func() {
		fmt.Println("正常退出")
	}()
	fmt.Println("嗨")

	defer func() {
		v := recover()
		fmt.Println("恐慌被恢复了", v)
	}()

	panic("拜拜")
	fmt.Println("执行不到这里")
}
