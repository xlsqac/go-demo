// package main
package main

import (
	"encoding/json"
	"fmt"
)

type Auther struct {
	Name string `json:"name,omitempty"`
	Age  int32  `json:"age,string,omitempty"`
}

func main() {
	jsonStr, err := json.Marshal(Auther{Name: "n"})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(jsonStr))
	a := &Auther{}
	err = json.Unmarshal(jsonStr, a)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(a)
}
