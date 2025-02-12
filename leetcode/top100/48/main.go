// package main
// https://leetcode.cn/problems/rotate-image/description/?envType=study-plan-v2&envId=top-100-liked
// 旋转图像
package main

import "fmt"

func main() {
	{
		matrix := [][]int{
			{1, 2, 3},
			{4, 5, 6},
			{7, 8, 9},
		}
		inPlaceRotate(matrix)
		fmt.Println(matrix)
	}
	{
		matrix := [][]int{
			{5, 1, 9, 11},
			{2, 4, 8, 10},
			{13, 3, 6, 7},
			{15, 14, 12, 16},
		}
		inPlaceRotate(matrix)
		fmt.Println(matrix)
	}
}

// rotate 旋转图像
// 时间复杂度 O(n^2)，需要遍历矩阵
// 空间复杂度 O(n^2)，需要一个新的矩阵
func rotate(matrix [][]int) {
	n := len(matrix)
	// 00, 01, 02 -> 02, 12, 22
	// 10, 11, 12 -> 01, 11, 21
	// 20, 21, 22 -> 00, 10, 20
	// 对于数字 1 来说，第 0 行的第 0 个元素，出现在倒数第 0 列的第 0 个元素
	// 对于数字 2 来说，第 0 行的第 1 个元素，出现在倒数第 0 列的第 1 个元素, 0列就是 n-1 列
	// 对于数字 3 来说，第 0 行的第 2 个元素，出现在倒数第 0 列的第 2 个元素，
	// 所以元素的新行就等于原来的列，i -> j
	// 第 0 行出现在倒数第 0 列(n-1)， 第一行出现在倒数第一列(n-1-1)
	// 第 2 行出现在倒数第二列也就是第 0 列(n-1-2)
	// 新的列就等于 n-1-i
	newMatrix := make([][]int, n)
	for i := 0; i < n; i++ {
		newMatrix[i] = make([]int, n)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			newMatrix[j][n-1-i] = matrix[i][j]
		}
	}
	copy(matrix, newMatrix)
}

// inPlaceRotate 旋转图像
// 原地修改，空间复杂度 O(1)
// 原地修改每次修改 4 个元素，所以不能全部遍历修改，这样会导致还原
func inPlaceRotate(matrix [][]int) {
	n := len(matrix)
	fmt.Println(n / 2)
	for i := 0; i < n/2; i++ {
		// n +1 / 2, n + 1 是处理 n 为奇数的时，中间一列不旋转的问题
		for j := 0; j < (n+1)/2; j++ {
			// tmp := matrix[i][j]
			// matrix[i][j] = matrix[n-j-1][i]
			// matrix[n-j-1][i] = matrix[n-i-1][n-j-1]
			// matrix[n-i-1][n-j-1] = matrix[j][n-i-1]
			// matrix[j][n-i-1] = tmp
			matrix[i][j], matrix[n-j-1][i], matrix[n-i-1][n-j-1], matrix[j][n-i-1] = matrix[n-j-1][i], matrix[n-i-1][n-j-1], matrix[j][n-i-1], matrix[i][j]

		}
	}
}
