package main

import (
    "fmt"
    "os"
)

func main() {
    var s, sep string
    fmt.Println(s)
    fmt.Println(sep)
    for i := 1; i < len(os.Args); i++ {
        s += sep + os.Args[i]
        sep = " "
    }
    fmt.Println(s)
}
