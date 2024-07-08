package main

import "fmt"

func change(s ...int) {
	s = append(s, 3)
	fmt.Println(s)
}

func change1(s []int) {
	s = append(s, 3)
	fmt.Println(s)
}

func main() {
	{
		slice := make([]int, 5)
		fmt.Println(slice, len(slice), cap(slice))
		slice = append(slice, 1)
		fmt.Println(slice, len(slice), cap(slice))

		slice1 := make([]int, 2, 5)
		fmt.Println(slice1, len(slice1), cap(slice1))
		slice1 = append(slice1, 1)
		fmt.Println(slice1, len(slice1), cap(slice1))
		fmt.Println("------------------------------")
	}

	{
		slice := make([]int, 5)
		slice[0] = 1
		slice[1] = 2
		change(slice...)
		fmt.Println(slice)
		change(slice[0:2]...)
		fmt.Println(slice)
		fmt.Println("------------------------------")
	}

	{
		slice := make([]int, 5)
		slice[0] = 1
		slice[1] = 2
		slice[2] = 4
		change1(slice)
		fmt.Println(slice)
		change1(slice[0:2])
		fmt.Println(slice)
	}
}
