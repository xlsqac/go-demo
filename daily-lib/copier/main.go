// package main
// 将一个结构体值复制到另一个类型的结构体值中
package main

import (
	"fmt"

	"github.com/jinzhu/copier"
)

// User
type User struct {
	Name string
	Age  int
}

// Employee
type Employee struct {
	Name string
	Age  int
	Role string
}

// main
func main() {
	user := User{Name: "dj", Age: 18}
	employee := Employee{}

	copier.Copy(&employee, &user)
	fmt.Printf("%#v\n", employee)
}
