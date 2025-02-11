// package main
// https://leetcode.cn/problems/set-matrix-zeroes/description/?envType=study-plan-v2&envId=top-100-liked
// 矩阵置零
package main

import (
	"fmt"
)

func main() {
	{
		matrix := [][]int{
			{1, 1, 1},
			{1, 0, 1},
			{1, 1, 1},
		}
		setZeroes(matrix)
		fmt.Println(matrix)
	}
	{
		matrix := [][]int{
			{0, 1, 2, 0},
			{3, 4, 5, 2},
			{1, 3, 1, 5},
		}
		setZeroes(matrix)
		fmt.Println(matrix)
	}
	{
		matrix := [][]int{
			{3, 7, 2, 8, 2},
			{2, 2, 4, 1, 8},
			{0, 5, 7, 6, 3},
			{8, 1, 0, 7, 7},
			{1, 3, 7, 4, 1},
			{6, 5, 5, 6, 3},
			{7, 1, 0, 1, 9},
			{5, 4, 4, 9, 7},
			{2, 2, 4, 1, 0},
			{7, 1, 1, 9, 1},
			{8, 0, 4, 7, 6},
			{7, 5, 1, 2, 3},
			{7, 2, 5, 9, 3},
		}
		setZeroes(matrix)
		fmt.Println(matrix)
	}
}

// setZeroes 矩阵置零
// 使用两个 map 记录需要置零的行和列，然后遍历 map 进行置零
// 时间复杂度 O(m*n)，遍历矩阵
func setZeroes(matrix [][]int) {
	rowM := make(map[int]bool)
	columnM := make(map[int]bool)
	rowLength := len(matrix)
	columnLength := 0
	for i := 0; i < rowLength; i++ {
		columnLength = len(matrix[i])
		for j := 0; j < columnLength; j++ {
			if matrix[i][j] == 0 {
				rowM[i] = true
				columnM[j] = true
			}
		}
	}

	if columnLength == 0 {
		return
	}

	for k := range rowM {
		for c := 0; c < columnLength; c++ {
			matrix[k][c] = 0
		}
	}
	for k := range columnM {
		for r := 0; r < rowLength; r++ {
			matrix[r][k] = 0
		}
	}

}
