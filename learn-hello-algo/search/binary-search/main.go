// package main
// https://www.hello-algo.com/chapter_searching/binary_search/
// 二分查找
// 给定一个长度为 的数组 nums，元素按从小到大的顺序排列且不重复。请查找并返回元素 target 在该数组中的索引

package main

import "fmt"

func main() {
	nums := []int{1, 3, 6, 8, 12, 15, 23, 26, 31, 35}
	target := 6
	index := binarcySearch(nums, target)
	fmt.Println(index)

}
func binarcySearch(nums []int, target int) int {
	i := 0
	j := len(nums) - 1
	for i <= j {
		m := (i + j) / 2
		if nums[m] < target {
			i = m + 1
		} else if nums[m] == target {
			return m
		} else {
			j = m - 1
		}
	}
	return -1
}
