package main

import "fmt"

func main() {
    var a = [5]int{1, 2, 3, 4, 5}
    var r [5]int

    // 自己描述：循环时修改 a，会影响 a 原本的值，但是不会影响循环时的值
    // 出题人描述：range 表达式是副本参与循环，修改的是 a 的值不是副本的值
    for i, v := range a {
        if i == 0 {
            a[1] = 12
            a[2] = 13
        }
        r[i] = v
    }
    fmt.Println("r = ", r) // {1, 2, 3, 4, 5}
    fmt.Println("a = ", a) // {1, 12, 13, 4, 5}

    // 想让 r 和 a 的输入一样
    var a1 = [5]int{1, 2, 3, 4, 5}
    var r1 [5]int

    // 用指针，会使原数组一直参与循环，每次取到的值都是修改后的值
    for i, v := range &a1 {
        if i == 0 {
            a1[1] = 12
            a1[2] = 13
        }
        r1[i] = v
    }

    fmt.Println("r1 = ", r1)
    fmt.Println("a1 = ", a1)
}
