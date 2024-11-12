// package main
// https://leetcode.cn/problems/best-time-to-buy-and-sell-stock/?envType=study-plan-v2&envId=top-100-liked
// 买卖股票的最佳时机
package main

import "fmt"

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit(prices))
}

// maxProfit 返回可以获取的最大利润
// 时间复杂度 O(n)，对切片循环一次
// 先把第一个元素当作最小值，然后遍历数组，如果当前元素比最小值小，则更新最小值
// 然后计算当前元素与最小值的差值，把每次的差值取最大值
func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	minPrice := prices[0]
	maxProfitPrice := 0
	for _, v := range prices[1:] {
		minPrice = min(v, minPrice)
		maxProfitPrice = max(v-minPrice, maxProfitPrice)

	}
	return maxProfitPrice
}
