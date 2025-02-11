// package main
// https://leetcode.cn/problems/rotate-array/description/?envType=study-plan-v2&envId=top-100-liked
// 轮转数组
package main

import (
	"fmt"
	"slices"
)

func main() {
	{
		nums := []int{1, 2, 3, 4, 5, 6, 7}
		k := 3
		rotate(nums, k)
		fmt.Println(nums)
	}
	{
		nums := []int{-1, -100, 3, 99}
		k := 2
		rotate(nums, k)
		fmt.Println(nums)
	}
}

// rotate 时间复杂度 O(n) 空间复杂度 O(n)
func rotate(nums []int, k int) {
	length := len(nums) // 7
	k = k % length      // 3
	// length -k = 4 nums[4:] = [5, 6, 7] nums[:4] = [1, 2, 3, 4]
	// append([5, 6, 7], [1, 2, 3, 4]) = [5, 6, 7, 1, 2, 3, 4]
	// copy 复制到一个切片中，在这里这个切片是原切片
	// 右移k位，就是末尾的 k 个元素移动到前面
	copy(nums, append(nums[length-k:], nums[:length-k]...))
}

// rotate1
// 三步反转法可以实现空间复杂度 O（1）
func rotate1(nums []int, k int) {
	length := len(nums)
	k = k % length
	slices.Reverse(nums)
	slices.Reverse(nums[:k])
	slices.Reverse(nums[k:])
}
