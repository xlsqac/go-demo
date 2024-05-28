// Package main
// https://leetcode.cn/problems/minimum-size-subarray-sum/description/
// 长度最小的子数组
// input: target = 7, nums = [2,3,1,2,4,3]
// return 2, 4+3
// 返回相加等于 target 的最小元素数量，2 + 3 + 1 + 2 = 4。 1 + 2 + 4 = 3，4 + 3 = 2
package main

import (
	"fmt"
	"math"
)

// minSubArrayLen O(N)
// left 慢指针，right 快指针，sum nums[left:right] 的和, res 最小子数组长度
// 如果 nums[right] + sum) >= target，说明当前窗口是满足条件的，需要不断缩小窗口以找到最小窗口 left++
// 不满足说明窗口不够大，right++ 扩大窗口
// 每次扩大窗口需要把新窗口的值加上，缩小窗口需要减去缩小的值
// 每次符合条件把 res 记录下来
func minSubArrayLen(target int, nums []int) int {
	left, right, sum := 0, 0, 0
	res := math.MaxInt32
	for right < len(nums) {
		if (nums[right] + sum) >= target {
			//fmt.Println(right, left, sum)
			res = min(res, right-left+1)
			sum -= nums[left]
			left++
		} else {
			sum += nums[right]
			right++
		}
	}
	if res == math.MaxInt32 {
		res = 0
	}
	//fmt.Println(subLens)
	return res
}

func main() {
	target := 7
	nums := []int{2, 3, 1, 2, 4, 3}
	res := minSubArrayLen(target, nums)
	fmt.Println(res)
}
