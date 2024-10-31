// package main
// https://leetcode.cn/problems/search-insert-position/description/?envType=study-plan-v2&envId=top-100-liked
// 搜索插入位置
package main

import (
	"fmt"
)

// main
func main() {
	nums := []int{1, 3, 5, 6}
	target := 1
	fmt.Println(searchInsert2(nums, target))
	target = 3
	fmt.Println(searchInsert2(nums, target))
	target = 5
	fmt.Println(searchInsert2(nums, target))
	target = 6
	fmt.Println(searchInsert2(nums, target))

	target = 2
	fmt.Println(searchInsert2(nums, target))

	target = 7
	fmt.Println(searchInsert2(nums, target))

	{
		nums := []int{1}
		target := 0
		fmt.Println(searchInsert2(nums, target))
	}
}

// searchInsert 搜索插入位置，二分查找，递归
// 时间复杂度 O(logn)
func searchInsert(nums []int, target int) int {
	index := len(nums) / 2
	if nums[index] == target {
		return index
	}
	if len(nums) == 1 {
		if nums[0] < target {
			return 1
		} else {
			return 0
		}
	}
	if len(nums) == 2 {
		if nums[1] < target {
			return 2
		} else if nums[0] < target {
			return 1
		}
		return 0
	}
	if nums[index] < target {
		return index + searchInsert(nums[index:], target)
	} else if nums[index] > target {
		return searchInsert(nums[:index], target)
	}
	return 0
}

// searchInsert2 搜索插入位置，二分查找，迭代
func searchInsert2(nums []int, target int) int {
	i, j := 0, len(nums)-1
	for i <= j {
		mid := i + (j-i)/2
		if nums[mid] < target {
			i = mid + 1
		} else if nums[mid] > target {
			j = mid - 1
		} else {
			return mid
		}
	}
	return i
	// i = 0; j = 3
	// mid = 0 + (3-0)/2 = 1
	// nums[1] = 3 target = 1
	// 3 > 1 j = 1 - 1 = 0

	// i = 0; j = 0
	// mid = 0 + (0-0)/2 = 0
	// num[0] = 1 target = 1
	// 1 == 1 return 0
	// 1 + (3-1)/2 = 2

}
