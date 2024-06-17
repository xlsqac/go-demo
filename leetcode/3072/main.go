// package main
// https://leetcode.cn/problems/distribute-elements-into-two-arrays-ii/description/
package main

import (
	"fmt"
)

// greaterCount 返回 arr 中 大于 val 的数量
func greaterCount(arr []int, val int) int {
	var ans int
	for _, v := range arr {
		if v > val {
			ans++
		}
	}
	return ans
}

func appendAndRecode(arr []int, val int, m map[int]int) []int {
	arr = append(arr, val)
	m[val-1]++
	return arr
}

func resultArray(nums []int) []int {
	arr1 := []int{nums[0]}
	arr2 := []int{nums[1]}
	map1 := map[int]int{nums[0] - 1: 1}
	map2 := map[int]int{nums[1] - 1: 1}

	for _, value := range nums[2:] {

		arr1Count := map1[value-1]
		arr2Count := map2[value-1]
		fmt.Println(map1, map2)
		fmt.Println(value, arr1Count, arr2Count)

		if arr1Count > arr2Count {
			arr1 = appendAndRecode(arr1, value, map1)
		} else if arr1Count < arr2Count {
			arr2 = appendAndRecode(arr2, value, map2)
		} else {
			if len(arr2) >= len(arr1) {
				arr1 = appendAndRecode(arr1, value, map1)
			} else {
				arr2 = appendAndRecode(arr2, value, map2)
			}
		}
	}
	ans := append(arr1, arr2...)

	return ans

}

func main() {
	nums := []int{2, 1, 3, 3}
	res := resultArray(nums)
	fmt.Println(res)

	nums = []int{5, 14, 3, 1, 2}
	res = resultArray(nums)
	fmt.Println(res)

	nums = []int{3, 3, 3, 3}
	res = resultArray(nums)
	fmt.Println(res)
}
