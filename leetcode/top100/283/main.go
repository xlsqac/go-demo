// package main
// https://leetcode.cn/problems/move-zeroes/description/?envType=study-plan-v2&envId=top-100-liked
// 移动零。给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。
// 必须在不复制数组的情况下原地对数组进行操作。
package main

import "fmt"

func main() {
	{
		nums := []int{0, 1, 0, 3, 12}
		moveZeroes1(nums)
		fmt.Println(nums)
	}

	{
		nums := []int{0}
		moveZeroes1(nums)
		fmt.Println(nums)
	}
	{
		nums := []int{1, 0}
		moveZeroes1(nums)
		fmt.Println(nums)
	}
	{
		nums := []int{1, 0, 0}
		moveZeroes1(nums)
		fmt.Println(nums)
	}

	{
		nums := []int{1, 2, 3, 1}
		moveZeroes1(nums)
		fmt.Println(nums)
	}
	{
		nums := []int{4, 2, 4, 0, 0, 3, 0, 5, 1, 0}
		moveZeroes1(nums)
		fmt.Println(nums)
	}
}

// 快慢指针
// 判断了慢指针的值，逻辑就显的不够清晰了
func moveZeroes(nums []int) {
	nLength := len(nums)

	// 小于 3 位的直接处理
	if nLength < 2 {
		return
	} else if nLength == 2 {
		if nums[0] == 0 && nums[1] != 0 {
			nums[0], nums[1] = nums[1], nums[0]
		}
		return
	}

	slow := 0
	fast := 1
	for fast < nLength && slow < nLength {
		if nums[slow] == 0 {
			if nums[fast] == 0 {
				fast++
			} else {
				nums[slow], nums[fast] = nums[fast], nums[slow]
				slow++
			}
		} else {
			slow++
			fast++
		}
	}
}

// 快慢指针都在 0，快指针每次都往前走
// 如果快指针不为 0，交换快慢指针的值，然后慢指针往前走一步
func moveZeroes1(nums []int) {
	slow := 0
	for fast := 0; fast < len(nums); fast++ {
		if nums[fast] != 0 {
			nums[slow], nums[fast] = nums[fast], nums[slow]
			slow++
		}
	}
}
