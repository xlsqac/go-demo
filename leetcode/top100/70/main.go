// package main
// https://leetcode.cn/problems/climbing-stairs/?envType=study-plan-v2&envId=top-100-liked
// 爬楼梯
package main

import "fmt"

func main() {
	n := 2
	fmt.Println(climbStairs(n))
	n = 3
	fmt.Println(climbStairs(n))
	n = 44
	fmt.Println(climbStairs1(n))
}

func climbStairs(n int) int {
	if n == 1 || n == 2 {
		return n
	}
	// 初始化 dp 表，用于存储子问题的解
	dp := make([]int, n+1)
	// 初始状态：预设最小子问题的解
	dp[1] = 1
	dp[2] = 2
	// 状态转移：从较小子问题逐步求解较大子问题
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

// climbStairs1 返回有多少种方法可以爬到楼顶
// 每一步的方法数是前 2 步的和
// 第一阶和第二阶的数量是固定的，1 和 2
// 所以从第三阶开始
func climbStairs1(n int) int {
	if n == 1 || n == 2 {
		return n
	}
	a, b := 1, 2
	for i := 3; i <= n; i++ {
		fmt.Println(i, a, b)
		a, b = b, a+b
	}
	return b
}

func backtrack(choices []int, state, n int, res []int) {
	// 如果到达了终点，就将结果数加1
	if state == n {
		res[0] = res[0] + 1
	}

	for _, choice := range choices {
		// 不允许超过第 n 阶
		if state+choice > n {
			continue
		}
		backtrack(choices, state+choice, n, res)
	}
}
