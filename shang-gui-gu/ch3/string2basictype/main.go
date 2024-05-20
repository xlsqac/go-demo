// string 转基本类型
// P50

package main

import (
	"fmt"
	"strconv"
)

func main() {
	var str1 string = "true"
	var str2 string = "12345678"
	var str3 string = "123.456"
	var str4 string = "hello"

	var b1 bool
	b1, _ = strconv.ParseBool(str1)
	fmt.Printf("b1 type: %T value: %v\n", b1, b1)

	var num1 int64
	num1, _ = strconv.ParseInt(str2, 10, 0)
	fmt.Printf("num1 type: %T value: %v\n", num1, num1)

	var num2 float64
	num2, _ = strconv.ParseFloat(str3, 64)
	fmt.Printf("num2 type: %T value: %v\n", num2, num2)

	var num3 int64
	num3, _ = strconv.ParseInt(str4, 10, 0) // hello 这种字符串会转成 0
	fmt.Printf("num3 type: %T value: %v\n", num3, num3)

}
