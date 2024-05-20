package main

import "fmt"

func main() {
	var d int = 97
	var weekDay = 7
	var week int = d / weekDay
	var day int = d % weekDay

	fmt.Printf("97天放假等于 %v 个星期零 %v 天\n", week, day)

	var f float32 = 134.2
	var c float32 = 5.0 / 9 * (f - 100)
	fmt.Printf("%v 华氏温度对应的设置温度是：%v\n", f, c)

}
