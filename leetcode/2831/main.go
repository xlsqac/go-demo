// package main
// https://leetcode.cn/problems/find-the-longest-equal-subarray/description/?envType=daily-question&envId=2024-05-23
// 找出最长等值子数组
// 滑动窗口
package main

import "fmt"

func longestEqualSubArray(nums []int, k int) int {
	// 默认是 0
	ans := 0
	pos := make(map[int][]int)
	for i, num := range nums {
		pos[num] = append(pos[num], i)

	}
	// 1 [0, 1, 4, 5]
	// 2 [2, 3]
	for _, vec := range pos {
		slow := 0
		// 能进循环说明最小有一个元素
		for fast := 0; fast < len(vec); fast++ {
			// 窗口的大小，nums[slow: fast] 这个区间一共有几个元素是 _
			windowSize := fast - slow
			numsCount := vec[fast] - vec[slow]
			// 0 - 0 - (0 - 0)= 0 > 2
			// 1 - 0 - (1 - 0) = 0 > 2
			// 4 - 0 - (2 - 0) = 2 > 2
			// 5 - 0 - (3 - 0) = 2 > 2
			// 3-0+1

			// numsCount-windowSize 当前窗口中非 _ 值的数量
			// numsCount 窗口在原数组中一共有几个元素
			// 如果中间非 _ 值的数量大于 K 说明删除后也不是完整的，需要慢指针往前走
			for numsCount-windowSize > k {
				slow++
			}
			ans = max(ans, windowSize+1)
		}
	}
	return ans

}

func main() {
	nums := []int{1, 1, 2, 2, 1, 1}
	k := 2

	subArrayLength := longestEqualSubArray(nums, k)
	fmt.Println(subArrayLength)
}
