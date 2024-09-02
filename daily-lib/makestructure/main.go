// package main
package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/mitchellh/mapstructure"
)

type Person struct {
	Name string
	Age  int
	Job  string
}

type Cat struct {
	Name  string
	Age   int
	Breed string
}

// main 把一个 json 转换成结构体，先把 json 转换成 map，再把 map 转换成结构体
func main() {
	datas := []string{`
    { 
      "type": "person",
      "name":"dj",
      "age":18,
      "job": "programmer"
    }
  `,
		`
    {
      "type": "cat",
      "name": "kitty",
      "age": 1,
      "breed": "Ragdoll"
    }
  `,
	}

	for _, data := range datas {
		var m map[string]interface{}
		err := json.Unmarshal([]byte(data), &m)
		if err != nil {
			log.Fatal(err)
		}
		switch m["type"].(string) {
		case "person":
			var p Person
			mapstructure.Decode(m, &p)
			fmt.Println("person", p)

		case "cat":
			var cat Cat
			mapstructure.Decode(m, &cat)
			fmt.Println("cat", cat)
		}
	}
}
