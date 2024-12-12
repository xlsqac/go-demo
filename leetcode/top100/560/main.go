// package main
// https://leetcode.cn/problems/subarray-sum-equals-k/description/?envType=study-plan-v2&envId=top-100-liked
// 和为 k 的子数组
package main

import "fmt"

func main() {
	{
		nums := []int{1, 1, 1}
		k := 2
		fmt.Println(2, subarraySum(nums, k))
	}
	{
		nums := []int{1, 2, 3}
		k := 3
		fmt.Println(2, subarraySum(nums, k))
	}
	{
		nums := []int{1, -1, 0}
		k := 0
		fmt.Println(3, subarraySum(nums, k))
	}
}

// subarraySum
// O(n) 利用前缀和 pre
// pre[i] = nums[0] + nums[1] + ... + nums[i]
// 前缀和可以是区间和

// pre 是当前元素的前缀和，如果 pre-k 存在于map 中，说明一个区间和等于 k
// 所有 pre-k 的次数就是子数组的个数
func subarraySum(nums []int, k int) int {
	ans, pre := 0, 0
	m := make(map[int]int)
	// 如果第一个元素正好等于 k，不这样写的话需要再循环判断，不然就少一次了
	m[0] = 1 // 默认 0 出现一次, 0 可以看做前缀和等于 k, 这样写可以在循环中少判断一次
	// nums 1 1 1 pre 1 2 3
	// pre =3 k = 2 pre-k =1
	for _, num := range nums {
		pre += num // 计算当前从 0 开始的前缀和
		// pre 是从 0 开始的前缀和，减去 k 就是区间和
		// 如果 pre-k 在 m 中出现过，说明存在区间和等于 k
		// 出现的次数，就等于子数组的次数
		// pre[i] - pre[j] 的值就等于 nums[j+1:i+1] 的和
		ans += m[pre-k]
		m[pre]++ // 保存每个前缀和的出现次数
	}

	return ans
}
