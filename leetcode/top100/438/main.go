// package main
// https://leetcode.cn/problems/find-all-anagrams-in-a-string/description/?envType=study-plan-v2&envId=top-100-liked
// 找到字符串中所有字母异位词
package main

import (
	"fmt"
)

func main() {
	{
		s := "cbaebabacd"
		p := "abc"
		fmt.Println(0, 6, findAnagrams(s, p))
	}
	{
		s := "abab"
		p := "ab"
		fmt.Println(0, 1, 2, findAnagrams(s, p))
	}
}

func findAnagrams(s string, p string) []int {
	// 存放每个窗口中每个字母的数量，一共 26个字母，从 a 开始
	countS := [26]int{}
	countP := [26]int{}
	ans := []int{}

	// 计算 p 中每个字母的数量
	for _, c := range p {
		countP[c-'a']++
	}

	for right, c := range s {
		// 窗口右端点进来
		countS[c-'a']++

		// 计算左断点的位置
		left := right - len(p) + 1
		// 刚开始的窗口长度小于 p 的长度，不需要比较
		if left < 0 {
			continue
		}

		// 比较两个数组
		if countS == countP {
			ans = append(ans, left)
		}
		// 左端点出去
		countS[s[left]-'a']--

	}
	return ans
}
