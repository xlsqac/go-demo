// package main
// https://leetcode.cn/problems/search-a-2d-matrix-ii/description/?envType=study-plan-v2&envId=top-100-liked
// 搜索二维矩阵
package main

import (
	"fmt"
	"sort"
)

func main() {
	matrix := [][]int{
		{1, 4, 7, 11, 15},
		{2, 5, 8, 12, 19},
		{3, 6, 9, 16, 22},
		{10, 13, 14, 17, 24},
		{18, 21, 23, 26, 30},
	}
	fmt.Println(searchMatrixFast(matrix, 5))
	fmt.Println(searchMatrixFast(matrix, 20))
}

// searchMatrix 搜索二维矩阵
// 时间复杂度 O(m*n) 暴力解法，直接遍历
func searchMatrix(matrix [][]int, target int) bool {
	row := len(matrix)
	column := len(matrix[0])
	for i := 0; i < row; i++ {
		for j := 0; j < column; j++ {
			if matrix[i][j] == target {
				return true
			}
		}
	}
	return false
}

// searchMatrixOptimized 搜索二维矩阵
// 利用有序的特性，发现比搜索元素大了就退出
func searchMatrixOptimized(matrix [][]int, target int) bool {
	row := len(matrix)
	column := len(matrix[0])
	for i := 0; i < row; i++ {
		for j := 0; j < column; j++ {
			if matrix[i][j] == target {
				return true
			}
			if matrix[i][j] > target {
				break
			}
		}
	}
	return false
}

// searchMatrixOptimized 搜索二维矩阵
// O(mlogn)，每次二分是 logn，一共 m 行
// 对每一行使用二分
func searchMatrixBinary(matrix [][]int, target int) bool {
	for _, row := range matrix {
		// 返回切片中大于等于 target 的最小索引
		// 如果元素大于所有，那么会返回切片长度
		i := sort.SearchInts(row, target)
		if i < len(row) && row[i] == target {
			return true
		}
	}
	return false
}

// searchMatrixFast	搜索二维矩阵
// O（m+n), m 和 n 是矩阵的行数和列数，每次循环是对 x或y 操作
// x 做多可以加 m 次，y 最多可以减 n 次
// 从右上角开始遍历，如果大于 target 就往左，如果小于 target 就往下
func searchMatrixFast(matrix [][]int, target int) bool {
	row, column := len(matrix), len(matrix[0])
	x, y := 0, column-1
	for x < row && y >= 0 {
		if matrix[x][y] == target {
			return true
		}
		if matrix[x][y] > target {
			y--
		} else {
			x++
		}
	}
	return false
}
