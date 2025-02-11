// package main
// 螺旋矩阵
// https://leetcode.cn/problems/spiral-matrix/?envType=study-plan-v2&envId=top-100-liked
package main

import "fmt"

func main() {
	{
		matrix := [][]int{
			{1, 2, 3},
			{4, 5, 6},
			{7, 8, 9},
		}
		// [1, 2, 3, 6, 9, 8, 7, 4, 5]
		// res[0]=matrix[0][0]
		// res[1]=matrix[0][1]
		// res[2]=matrix[0][2]

		// res[3]=matrix[1][2]
		// res[4]=matrix[2][2]

		// res[5]=matrix[2][1]
		// res[6]=matrix[2][0]

		// res[7]=matrix[1][0]
		// res[8]=matrix[1][1]
		res := spiralOrder(matrix)
		fmt.Println(res)
	}
	{
		matrix := [][]int{
			{1, 2, 3, 4},
			{5, 6, 7, 8},
			{9, 10, 11, 12},
		}
		res := spiralOrder(matrix)
		fmt.Println(res)
	}
}

// spirlOrder 螺旋矩阵
// 时间复杂度 O(m*n)
// 空间复杂度 O(m*n)
func spiralOrder(matrix [][]int) []int {
	// 矩阵行的长度和列的长度
	rows := len(matrix)
	columns := len(matrix[0])
	if rows == 0 || columns == 0 {
		return []int{}
	}
	// 辅助矩阵，判断路径是否进入之前访问过的位置
	visited := make([][]bool, rows)
	for i := 0; i < rows; i++ {
		visited[i] = make([]bool, columns)
	}
	// 4个方向，向右，向下，向左，向上
	// 向右就是行不变，列+1
	// 向下就是列不变，行+1
	// 向左就是行不变，列-1
	// 向上就是列不变，行-1
	total := rows * columns
	res := make([]int, total)
	// 起始坐标
	row, column := 0, 0
	directions := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	// 起始方向
	directionIndex := 0

	for i := 0; i < total; i++ {
		res[i] = matrix[row][column]
		// 走过的进行标记
		visited[row][column] = true
		// 根据现在前进的方向，确定下一个位置
		nextRow, nextCloumn := row+directions[directionIndex][0], column+directions[directionIndex][1]
		// 判断下一个位置是否越界，如果越界或者已经访问过就需要更换方向
		if nextRow < 0 || nextRow >= rows || nextCloumn < 0 || nextCloumn >= columns || visited[nextRow][nextCloumn] {
			directionIndex = (directionIndex + 1) % 4
		}
		// 根据现在前进的方向，确定下一个位置
		row = row + directions[directionIndex][0]
		column = column + directions[directionIndex][1]
	}

	return res
}
