// package main
// https://leetcode.cn/problems/find-missing-observations/description/?envType=daily-question&envId=2024-05-27
package main

import (
	"fmt"
)

// missingRolls
// 暴力求值
// 先求出整个数组和 rolls 的差值
// 根据差值和 n 求出确实的数据中的平均数，以及余数部分
// 把平均数加入到 ans 中，余数一定小于 n, 所以从前往后给小于余数的索引加 1 即可
func missingRolls(rolls []int, mean int, n int) []int {
	const (
		MAXROLL = 6
	)
	length := len(rolls)
	sum := mean * (n + length)
	//fmt.Println("[rolls sum]", sum)
	sumRolls := 0
	for _, roll := range rolls {
		sumRolls += roll
	}
	//fmt.Println("[rolls m sum]", sumRolls)
	missingSum := sum - sumRolls
	//fmt.Println("[rolls missing sum]", missingSum)
	// roll min 1 max 6
	// missingSum < n*1 或 missingSum > n*6 说明没有满足条件的情况
	if missingSum < n || missingSum > n*MAXROLL {
		return nil
	}
	//
	avg, remainder := missingSum/n, missingSum%n
	//fmt.Println("[rolls missing avg]", missingSum, n, avg)
	ans := make([]int, n)
	for i := range n {
		ans[i] = avg
		if i < remainder {
			ans[i]++
		}
	}
	return ans
}

func main() {
	rolls := []int{3, 2, 4, 3}
	mean, n := 4, 2
	missRolls := missingRolls(rolls, mean, n)
	fmt.Println(missRolls)
	fmt.Println()

	rolls = []int{1, 5, 6}
	mean, n = 3, 4
	missRolls = missingRolls(rolls, mean, n)
	fmt.Println(missRolls)
	fmt.Println()

	rolls = []int{1}
	mean, n = 3, 1
	missRolls = missingRolls(rolls, mean, n)
	fmt.Println(missRolls)
	fmt.Println()

	rolls = []int{1, 2, 3, 4}
	mean, n = 6, 4
	missRolls = missingRolls(rolls, mean, n)
	fmt.Println(missRolls)
	fmt.Println()

	rolls = []int{4, 5, 6, 2, 3, 6, 5, 4, 6, 4, 5, 1, 6, 3, 1, 4, 5, 5, 3, 2, 3, 5, 3, 2, 1, 5, 4, 3, 5, 1, 5}
	mean, n = 4, 40
	missRolls = missingRolls(rolls, mean, n)
	fmt.Println(missRolls)
	fmt.Println()

}
