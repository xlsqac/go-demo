package main

import (
    "fmt"
    "os"
    "strings"
    "time"
)

func echoFor(args []string) {
    start := time.Now()
    s, sep := "", ""
    for i := 1; i < len(args); i++ {
        s += sep + args[i]
        sep = " "
    }
    fmt.Println(s)
    secs := time.Since(start).Seconds()
    fmt.Printf("for: %f\n", secs)
}

func echoJoin(args []string) {
    start := time.Now()
    fmt.Println(strings.Join(args[1:], " "))
    secs := time.Since(start).Seconds()
    fmt.Printf("join: %f\n", secs)
}

func main() {
    echoFor(os.Args)
    echoJoin(os.Args)
}
