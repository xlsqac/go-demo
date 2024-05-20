package main

import (
    "fmt"
    "os"
)

func main() {
    for index, arg := range os.Args {
        fmt.Printf("index: %d value: %s\n", index, arg)
    }
}
