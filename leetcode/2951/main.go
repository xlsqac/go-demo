// package main
// https://leetcode.cn/problems/find-the-peaks/description/
// [2,5,4] 返回数组中峰值的索引
// 峰值：严格大于相邻元素，第一个元素和最后一个元素不是峰值
// O(n)
package main

import "fmt"

// 遍历一次，大于相邻元素就返回
func findPeaks(mountain []int) []int {
	var ans []int
	length := len(mountain)

	for i := 1; i < length-1; i++ {
		if mountain[i] > mountain[i-1] && mountain[i] > mountain[i+1] {
			ans = append(ans, i)
		}
	}

	return ans
}

func main() {
	mountain := []int{2, 4, 4}
	peaks := findPeaks(mountain)
	fmt.Println(peaks)
	fmt.Println()

	mountain = []int{1, 4, 3, 8, 5}
	fmt.Println(mountain[1:4])
	peaks = findPeaks(mountain)
	fmt.Println(peaks)
}
