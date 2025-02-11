// package main
// https://leetcode.cn/problems/product-of-array-except-self/?envType=study-plan-v2&envId=top-100-liked
// 除自身以外数组的乘积
package main

import (
	"fmt"
)

func main() {
	{
		nums := []int{1, 2, 3, 4}
		fmt.Println(productExceptSelf(nums))
	}
}

// productExceptSelf 除自身以外数组的乘积
// 时间复杂度O(n) 空间复杂度O(n)
// 先求出从左到右每个前缀的乘积，再求出从右到左每个前缀的乘积，最后将两个数组相乘
// 没有前缀时，乘积为1
// 最后的结果，每个数除自身外的乘积等于前缀乘积乘以后缀乘积
func productExceptSelf(nums []int) []int {
	length := len(nums)
	lProduct, rProduct, answer := make([]int, length), make([]int, length), make([]int, length)
	lProduct[0], rProduct[length-1] = 1, 1
	for i := 1; i < length; i++ {
		lProduct[i] = lProduct[i-1] * nums[i-1]
	}
	fmt.Println(lProduct)
	for i := length - 2; i >= 0; i-- {
		rProduct[i] = rProduct[i+1] * nums[i+1]
	}
	fmt.Println(rProduct)
	for i := 0; i < length; i++ {
		answer[i] = lProduct[i] * rProduct[i]
	}
	return answer
}
