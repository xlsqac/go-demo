// 接收键盘输入
// fmt.Scanf
package main

import "fmt"

func main() {

    var name string
    var age byte
    var sal float32
    var isPass bool

    fmt.Println("请输入你的姓名、年龄、薪水和是否通过考试并使用空格隔开。")
    fmt.Scanf("%s %d %f %t", &name, &age, &sal, &isPass)
    fmt.Printf("名字是：%v\n年龄是：%v\n薪水是：%v\n是否通过考试：%v\n", name, age, sal, isPass)

}
