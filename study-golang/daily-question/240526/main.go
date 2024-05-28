package main

func add(args ...int) int {
	sum := 0
	for _, arg := range args {
		sum += arg
	}
	return sum
}

func main() {
	s := []int{1, 2}
	add(1, 2, 7)
	// add([]int{1, 2})
	add([]int{1, 2}...)
	add(s...)

}
