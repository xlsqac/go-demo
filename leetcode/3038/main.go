// package main
// https://leetcode.cn/problems/maximum-number-of-operations-with-the-same-score-i/
package main

import "fmt"

// maxOperations O(n)
// 先求出分数，然后把剩下的元素 /2 算出需要循环的次数
// 每次循环判断前 2 个元素和是否和分数相等，不相等直接退出循环
func maxOperations(nums []int) int {
	ans := 0
	if len(nums) < 2 {
		return ans
	}
	ans++
	score := nums[0] + nums[1]
	loopTimes := len(nums[2:]) / 2
	//fmt.Println(loopTimes)
	index := 2
	for i := 0; i < loopTimes; i++ {
		sum := nums[index] + nums[index+1]
		if sum != score {
			break
		}
		ans++
		index += 2

	}
	return ans

}

func main() {
	var nums []int
	var res int
	nums = []int{3, 2, 1, 4, 5}
	res = maxOperations(nums)
	fmt.Println(res)

	nums = []int{3, 2, 6, 1, 4}
	res = maxOperations(nums)
	fmt.Println(res)

	nums = []int{2, 5, 3, 4, 5, 2, 6, 3, 7}
	res = maxOperations(nums)
	fmt.Println(res)

	nums = []int{1, 1, 1, 1, 1, 1}
	res = maxOperations(nums)
	fmt.Println(res)
}
