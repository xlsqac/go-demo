package main

import (
	"encoding/json"
	"fmt"
	"time"
)

//type Person struct {
//	name  string
//	hobby string
//}
//
//func (p Person) MarshalJSON() ([]byte, error) {
//	return []byte(`{"name": "` + p.name + `", "hobby": "` + p.hobby + `"}`), nil
//}

func main() {
	i := 0b1111
	t := struct {
		time.Time
		N int
	}{
		time.Date(2020, 12, 20, 0, 0, 0, 0, time.UTC),
		5,
	}

	m, _ := json.Marshal(t)
	fmt.Printf("%s\n", m)
	fmt.Printf("%d\n", i)

	//person := Person{name: "name", hobby: "golang"}
	//m1, _ := json.Marshal(person)
	//fmt.Printf("%s\n", m1)
}
