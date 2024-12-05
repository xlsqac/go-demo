// package main
// https://leetcode.cn/problems/container-with-most-water/description/?envType=study-plan-v2&envId=top-100-liked
// 容纳最多水的容器
package main

import "fmt"

func main() {
	{
		nums := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
		fmt.Println(49, maxArea(nums))
	}
	{
		nums := []int{1, 1}
		fmt.Println(1, maxArea(nums))
	}
	{
		nums := []int{1, 2, 1}
		fmt.Println(2, maxArea(nums))
	}
	{
		nums := []int{1, 2, 4, 3}
		fmt.Println(4, maxArea(nums))
	}
	{
		nums := []int{4, 3, 2, 1, 4}
		fmt.Println(4, maxArea(nums))
	}
}

// maxArea return a int value 返回最大容器值
// 时间复杂度 n
// 双指针，左指针从开头往后移动，右指针从后往前移动
// 如果左指针的值小于右指针的值，左指针往后移动，否则右指针往前移动
// 计算每次指针区域内的值，最后返回最大值即可
// 因为容器不能倾斜，所以区域的最小值决定了容器的大小
func maxArea(height []int) int {
	lenHeight := len(height)
	left, right := 0, lenHeight-1
	ans := 0
	for left < right {
		ans = max(ans, min(height[left], height[right])*(right-left))
		if height[right] > height[left] {
			left++
		} else {
			right--
		}

	}
	return ans
}
