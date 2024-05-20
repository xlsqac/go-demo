/*
交换两个变量的值
*/
package main

import "fmt"

func temp() {
    //临时变量法
    fmt.Println("使用临时变量交换值")

    var a int = 1
    var b int = 2
    fmt.Printf("交换前：a = %v；b = %v。\n", a, b)

    var temp int = a
    a = b
    b = temp
    fmt.Printf("交换后：a = %v；b = %v。\n", a, b)
}

func calculate() {
    fmt.Println("使用计算法交换值")

    var a int = 1
    var b int = 2
    fmt.Printf("交换前：a = %v；b = %v。\n", a, b)

    a = a + b
    b = a - b // a + b - b
    a = a - b // a + b - a
    fmt.Printf("交换后：a = %v；b = %v。\n", a, b)
}

func main() {
    temp()
    fmt.Printf("")
    calculate()

}
