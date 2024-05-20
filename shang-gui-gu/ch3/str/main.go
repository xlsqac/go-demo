package str

import "fmt"

var s = "test"

func main() {
	fmt.Println("s =", s)
	fmt.Printf("s[0] = %c\n", s[0]) // 可以 s[0] 取字符串中的某个字符，但是不能修改单个字符 eg: s[0] = "b" error
}
