package main

import "fmt"

func main() {
	var s = make(map[string]int)
	delete(s, "h")
	value := s["h"]
	fmt.Println(value)

}
