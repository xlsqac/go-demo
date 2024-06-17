// package main
// https://leetcode.cn/problems/separate-black-and-white-balls/
package main

import "fmt"

const (
	WHITE = iota + 48
	BLACK
)

// minimumSteps O(n) 双指针，碰到白球计算左右指针的差值，然后左指针往前走一步
// 碰到白球左边指针前进，左右指针的差值就是中间黑球的个数也就是白球移动所需的步数
func minimumSteps(s string) int64 {
	ans := int64(0)
	left, right := 0, 0

	for ; right < len(s); right++ {
		if s[right] == '0' {
			ans += int64(right - left)
			left++
		}

	}

	return ans
}

func main() {
	s := "101"
	res := minimumSteps(s)
	fmt.Println(res)

	s = "100"
	res = minimumSteps(s)
	fmt.Println(res)

	s = "0001"
	res = minimumSteps(s)
	fmt.Println(res)
}
