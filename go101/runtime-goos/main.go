// package main
package main

import "runtime"
import "fmt"

func main() {
	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)

}
