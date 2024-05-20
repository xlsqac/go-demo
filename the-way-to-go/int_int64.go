package main

import "fmt"

var a int64 = 1
var b int64 = 2


func main() {
    b = b + a
    fmt.Printf("b is %d\n", b)
}

