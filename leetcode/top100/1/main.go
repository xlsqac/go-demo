// package main
// https://leetcode.cn/problems/two-sum/description/?envType=study-plan-v2&envId=top-100-liked
// 两数之和，给定一个切片和一个目标值，返回两个数的下标，两个数的和等于目标值
package main

import (
	"fmt"
)

func main() {
	{
		nums := []int{2, 7, 11, 15}
		target := 9
		res := twoSum(nums, target)
		fmt.Println(res, []int{0, 1})
	}
	{
		nums := []int{3, 2, 4}
		target := 6
		res := twoSum(nums, target)
		fmt.Println(res, []int{1, 2})
	}
	{
		nums := []int{3, 3}
		target := 6
		res := twoSum(nums, target)
		fmt.Println(res, []int{0, 1})
	}
}

// 利用哈希表，判断每次的差是否在 hash 中, O(n)
func twoSum(nums []int, target int) []int {
	hashTable := make(map[int]int)

	for i, v := range nums {
		// 判断目标值减去当前值的差是否在 hash 中
		if pos, ok := hashTable[target-v]; ok {
			return []int{pos, i}
		}
		// 不在的话把当前值加入到 hash 中
		hashTable[v] = i
	}
	return []int{}
}

// 暴力枚举 O(n^2)
func twoSum1(nums []int, target int) []int {
	var output []int
	lentch := len(nums)
	if lentch < 2 {
		return output
	}

	for i, v := range nums {
		for j := i + 1; j < lentch; j++ {
			if v+nums[j] == target {
				output = append(output, i, j)
			}
		}
	}
	return output
}
