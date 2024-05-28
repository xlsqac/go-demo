package main

import "fmt"

func removeElement(nums []int, val int) int {
	slow := 0
	for fast := 0; fast < len(nums); fast++ {
		if nums[fast] != val {
			nums[slow] = nums[fast]
			slow++
		}
	}
	return slow

}

func main() {
	nums := []int{3, 2, 2, 3}
	val := 3
	length := removeElement(nums, val)
	fmt.Println(nums[:length])
}
