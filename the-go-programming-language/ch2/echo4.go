package main

import (
	"flag"
	"fmt"
	"strings"
)

var n = flag.Bool("n", false, "忽略行尾的换行符")
var sep = flag.String("s", " ", "分隔字符")

func main() {
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}
}
