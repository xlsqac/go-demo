// package main
// https://leetcode.cn/problems/pascals-triangle/solutions/510638/yang-hui-san-jiao-by-leetcode-solution-lew9/?envType=study-plan-v2&envId=top-100-liked
// 杨辉三角
package main

import "fmt"

func main() {
	fmt.Println(generate(5))
}

// generate
// 从上一层推下一层
// 先把第一层放好，循环中对下一层处理
// 下一层的 0 和 -1 都是 1，中间的元素是上一层的 j-1 和 j 之和
func generate(numRows int) [][]int {
	// 生成一个二维数组，外层长度是 numRows
	nums := make([][]int, numRows)
	// 给第一层赋值
	nums[0] = []int{1}

	// 循环数组
	for i, row := range nums {
		// 因为是给下一层赋值，所以到最后一层的时候直接返回
		if i == numRows-1 {
			return nums
		}
		// 下一层的数组，先把 0 和 -1 赋值为 1
		nextRow := make([]int, i+2)
		nextRow[0] = 1
		nextRow[i+1] = 1
		// 给中间添加元素
		for j := 1; j < i+1; j++ {
			nextRow[j] = row[j-1] + row[j]
		}
		nums[i+1] = nextRow
	}
	return nums
}
