// package main
// https://leetcode.cn/problems/single-number/description/?envType=study-plan-v2&envId=top-100-liked
// 只出现一次的数字
package main

import "fmt"

func main() {
	{
		nums := []int{2, 2, 1}
		fmt.Println(singleNumber(nums))
	}

	{
		nums := []int{4, 1, 2, 1, 2}
		fmt.Println(singleNumber(nums))
	}

	{
		nums := []int{1}
		fmt.Println(singleNumber(nums))
	}
}

func singleNumber(nums []int) int {
	single := 0
	for _, num := range nums {
		single ^= num
	}
	return single
}
