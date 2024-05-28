// package main
package main

import (
	"fmt"
)

const Size = 32

const a uint = 1

type Person struct {
	name string
	age  int
}

func main() {
	//[5]string
	//[Size]int
	//[Size][]byte
	//[Size]Person
	//[4]bool{true, false, true, false}
	//b := [4]bool{0: false, 2: true, 1: true, 3: false}
	b := [4]bool{1: false, true, true}
	fmt.Println(b)

	n := [3]int{a: 1}
	fmt.Println(n)

	pa := &[...]bool{false, true, true, false}
	fmt.Printf("%T\n", pa) // *[4]

	var array [16]int
	var array1 [16]int
	array1[15] = 15
	fmt.Println(array1[15])
	fmt.Println(array == array1)

	var ar [5]int
	fmt.Println(len(ar), cap(ar))
	var s []int
	fmt.Println(len(s), cap(s)) // 0 0

	array2 := [5]int{2, 3, 5, 7}
	ps1 := &array2[1]
	fmt.Println(*ps1)

	array2[1] = 99
	fmt.Println(*ps1)

	for key, element := range array2 {
		array2[key] = element + 1
		fmt.Println(array2)
	}
	fmt.Println(array2)

	persons := [2]Person{{"Alice", 28}, {"Bob", 25}}
	for i, p := range persons {
		fmt.Println(i, p)
		persons[i].name = "jack"
		p.age = 31
	}
	fmt.Println(persons)

}
