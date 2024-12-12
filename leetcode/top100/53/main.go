// package main
// https://leetcode.cn/problems/maximum-subarray/description/?envType=study-plan-v2&envId=top-100-liked
// 最大子数组和
package main

import "fmt"

func main() {
	{
		nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
		fmt.Println(6, maxSubArray2(nums))
	}

	{
		nums := []int{1}
		fmt.Println(1, maxSubArray2(nums))
	}
	{
		nums := []int{5, 4, -1, 7, 8}
		fmt.Println(23, maxSubArray2(nums))
	}
	{
		nums := []int{-1, 0, -2}
		fmt.Println(0, maxSubArray2(nums))
	}
}

// maxSubArray 返回最大和
// N^2，暴力做法，双层循环，计算每个子数组的和，取最大值
func maxSubArray(nums []int) int {
	ans := nums[0]
	for i := 0; i < len(nums); i++ {
		sum := nums[i]
		ans = max(ans, sum)
		for j := i + 1; j < len(nums); j++ {
			sum += nums[j]
			ans = max(ans, sum)
		}
	}
	return ans
}

// maxSubArray2
// O(n)，贪心
// 先计算 1 个元素的和，之后循环判断当前之前和是否小于 0，小于 0 则丢弃之前的和，从当前元素开始计算
// 因为负数相加不如直接取最大的元素
func maxSubArray2(nums []int) int {
	maxSum := nums[0]
	sum := nums[0]
	for _, num := range nums[1:] {
		if sum < 0 {
			sum = num
		} else {
			sum += num
		}
		maxSum = max(sum, maxSum)
	}
	return maxSum
}

func prefixSum(arr []int) []int {
	n := len(arr)
	prefix := make([]int, n)
	prefix[0] = arr[0] // 第一个元素的前缀和即为它本身

	for i := 1; i < n; i++ {
		prefix[i] = prefix[i-1] + arr[i] // 当前前缀和是前一个前缀和加上当前元素
	}

	return prefix
}
