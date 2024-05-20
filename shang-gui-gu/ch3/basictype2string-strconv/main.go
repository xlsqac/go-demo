// 基本数据类型转 string 使用 strconv 包
// P49
package main

import (
	"fmt"
	"strconv"
)

func main() {
	var num3 int = 99
	var num4 float64 = 23.456
	var num5 int = 4567
	var b2 bool = true

	str := strconv.FormatInt(int64(num3), 10)
	fmt.Printf("str type: %T value:%q\n", str, str)

	str = strconv.FormatFloat(num4, 'f', 10, 64) // 保留的位数小于实际位数时会进行四舍五入
	fmt.Printf("str type: %T value:%q\n", str, str)

	str = strconv.Itoa(num5) // 等价于 FormatInt(int64(num5), 10)
	fmt.Printf("str type: %T value:%q\n", str, str)

	str = strconv.FormatBool(b2)
	fmt.Printf("str type: %T value:%q\n", str, str)
}
