// package main
// https://leetcode.cn/problems/majority-element/solutions/146074/duo-shu-yuan-su-by-leetcode-solution/?envType=study-plan-v2&envId=top-100-liked
// 多数元素
package main

import "fmt"

func main() {
	nums := []int{3, 2, 3}
	fmt.Println(majorityElement2(nums))

	{
		nums := []int{2, 2, 1, 1, 1, 2, 2}
		fmt.Println(majorityElement2(nums))
	}

	{
		nums := []int{3, 3, 4}
		fmt.Println(majorityElement2(nums))
	}
}

// majorityElement 哈希表统计
func majorityElement(nums []int) int {
	elementToCount := make(map[int]int)
	maxCountNum := nums[0]
	maxCount := 0
	for _, v := range nums[1:] {
		count := elementToCount[v] + 1
		elementToCount[v] = count
		if count > maxCount {
			maxCount = count
			maxCountNum = v
		}
	}
	// fmt.Println(elementToCount)
	return maxCountNum
}

// majorityElement2 Boyer-Moore 投票
// 先假定第一个数是多数，然后遍历数组，遇到相同的数+1，遇到不同的数-1，当计数器为0时，假定下一个数是多数，重复上述过程
func majorityElement2(nums []int) int {
	res := nums[0]
	counter := 1
	for _, v := range nums[1:] {
		if v != res {
			if counter--; counter == 0 {
				res = v
				counter = 1
			}
		} else {
			counter++
		}
	}
	return res
}
