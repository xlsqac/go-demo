// package main
// https://leetcode.cn/problems/find-missing-and-repeated-values/
package main

import (
	"fmt"
)

// 先遍历 grid 把每个数字出现的次数统计出来，把次数为 2 的放到 ans[0] 中
// 出现的元素范围是 [1: n*n]，遍历该范围并且和 m 进行比较找到出现次数为 0 的元素退出。
// O(n * n)
// 也可以用数组统计，索引就是 map 的键
func findMissingAndRepeatedValues(grid [][]int) []int {
	gLen := len(grid)
	maxNum := gLen * gLen
	m := make(map[int]int)
	ans := []int{0, 0}

	for _, v := range grid {
		for _, w := range v {
			m[w] += 1
			if m[w] == 2 {
				ans[0] = w
			}
		}
	}

	// missing
	for i := 1; i <= maxNum; i++ {
		if m[i] == 0 {
			ans[1] = i
			return ans
		}
	}
	return ans

}

func main() {
	grid := [][]int{{1, 3}, {2, 2}}
	ans := findMissingAndRepeatedValues(grid)
	fmt.Println(ans)

	grid = [][]int{{9, 1, 7}, {8, 9, 2}, {3, 4, 6}}
	ans = findMissingAndRepeatedValues(grid)
	fmt.Println(ans)
}
