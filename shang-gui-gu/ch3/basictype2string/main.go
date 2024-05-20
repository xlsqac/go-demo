// 基本数据类型转 string
// P49
package main

import "fmt"

func main() {
	var num1 int = 99
	var num2 float64 = 23.456
	var b bool = true
	var myChar byte = 'h'
	var str string

	str = fmt.Sprintf("%d", num1)
	fmt.Println("int to str, str = ", str)

	str = fmt.Sprintf("%f", num2)
	fmt.Println("float64 to str, str = ", str)

	str = fmt.Sprintf("%t", b)
	fmt.Println("bool to str, str = ", str)

	str = fmt.Sprintf("%c", myChar)
	fmt.Println("byte to str, str = ", str)
}
