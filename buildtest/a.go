package main

import (
	"fmt"
	"os"
)

func main() {
	_, err := os.Stat("a_test\uFE0E.go")
	if os.IsNotExist(err) {
		fmt.Println("File does not exist")
	} else if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("File exists")
	}

	{
		_, err := os.Stat("a_test.go")
		if os.IsNotExist(err) {
			fmt.Println("File does not exist")
		} else if err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Println("File exists")
		}
	}
}
