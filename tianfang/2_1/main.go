package main
import "fmt"

const (
    a = iota
    b = iota
)

const (
    name = "heelo"
    name1 = "11"
    c = iota
    d = iota
)

func main() {
    fmt.Println(a)
    fmt.Println(b)
    fmt.Println(c)
    fmt.Println(d)
}
