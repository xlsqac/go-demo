package main

import "fmt"

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	} else if len(nums) == 1 {
		return 1
	}

	slow := 0
	for fast := 1; fast < len(nums); fast++ {
		if nums[slow] != nums[fast] {
			nums[slow] = nums[fast]
			slow++
		}
	}
	return slow + 1
}

func main() {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	times := removeDuplicates(nums)
	fmt.Println(times)
}
