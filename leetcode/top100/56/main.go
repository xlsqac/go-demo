// package main
// https://leetcode.cn/problems/merge-intervals/?envType=study-plan-v2&envId=top-100-liked
// 合并区间
package main

import (
	"fmt"
	"sort"
)

func main() {
	{
		intervals := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
		// [[1,6],[8,10],[15,18]]
		fmt.Println(merge(intervals))
	}
	{
		intervals := [][]int{{1, 4}, {4, 5}}
		// [[1,5]]
		fmt.Println(merge(intervals))
	}
	{
		s := []string{"t", "e", "s", "t"}
		s[0] = s[1]
		fmt.Println(s)
	}

}

// merge 合并区间
// 时间复杂度 O(nlogn)，排序 nlogn，遍历 n，总体是 nlogn
// 先对数组排序，然后把 0 号元素加入结果集，然后遍历数组
// 如果右边的数组的 0 号元素小于等于结果集的最后一个元素的 1 号元素，说明有重合区间
// 然后把重合区间整合好即可
func merge(intervals [][]int) [][]int {
	ans := [][]int{}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	ans = append(ans, intervals[0])
	for _, interval := range intervals[1:] {
		ansIndex := len(ans) - 1
		// 有重合区间
		if interval[0] <= ans[ansIndex][1] {
			if interval[1] > ans[ansIndex][1] {
				ans[ansIndex][1] = interval[1]
			}
		} else {
			ans = append(ans, interval)
		}
	}
	// fmt.Println(intervals)
	return ans
}
