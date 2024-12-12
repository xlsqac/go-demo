// package main
// https://leetcode.cn/problems/longest-consecutive-sequence/description/?envType=study-plan-v2&envId=top-100-liked
// 最长连续序列，在一个未排序的数组中，找出最长的连续序列的长度，[100,4,200,1,3,2] return 4
package main

import (
	"fmt"
	"sort"
)

func main() {
	{
		nums := []int{100, 4, 200, 1, 3, 2}
		fmt.Println(4, longestConsecutiveMap(nums))
	}

	{
		nums := []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}
		fmt.Println(4, longestConsecutiveMap(nums))
	}
	{
		nums := []int{1, 2, 0, 1}
		fmt.Println(4, longestConsecutiveMap(nums))
	}
}

// longestConsecutive return int 最长连续序列的长度
// 时间复杂度 nlogn，sort.Ints nlogn，遍历数组 n，总体 nlogn
// 排序后，遍历数组，记录最长连续序列的长度
func longestConsecutive(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}

	sort.Ints(nums)
	prevNum := nums[0]
	longestConsecutiveNum := 1
	maxLongestConsecutiveNum := 1
	for _, num := range nums[1:] {
		if num == prevNum+1 {
			longestConsecutiveNum++
			maxLongestConsecutiveNum = max(maxLongestConsecutiveNum, longestConsecutiveNum)
		} else if num == prevNum {
			continue
		} else {
			longestConsecutiveNum = 1
		}
		prevNum = num

	}
	return maxLongestConsecutiveNum
}

// longestConsecutiveMap return int 返回最长序列的长度
// 时间复杂度 n，遍历数组 n
// 先遍历数组把数组存在 map 中
// 然后遍历 map，如果 num-1 不在 map 中，那么 num 就是当前序列的起点，遍历序列，记录最长序列的长度
// 如果当前序列终端就再次寻找遍历的起点，最后返回最长的序列长度
func longestConsecutiveMap(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}

	// 把数字添加到 map 中，实际是一个 set
	numMap := make(map[int]struct{})
	for _, num := range nums {
		numMap[num] = struct{}{}
	}

	// 最长序列长度
	maxLongestConsecutiveNum := 1

	// 遍历 map
	for num := range numMap {
		// 从 map 获取 num -1 的 key，如果这个 key 不存在，说明这个 num 是序列的起点
		if _, ok := numMap[num-1]; !ok {
			startNum := num
			longestConsecutiveNum := 1

			// 以起点开始从 map 中一个一个找，找不到就说明这个序列结束了
			for {
				if _, ok := numMap[startNum+1]; ok {
					startNum++
					longestConsecutiveNum++
				} else {
					break
				}
			}

			// 每次序列结束都比较一次，取最长的序列
			maxLongestConsecutiveNum = max(maxLongestConsecutiveNum, longestConsecutiveNum)

		}
	}
	return maxLongestConsecutiveNum
}
