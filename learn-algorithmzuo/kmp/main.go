// package main
package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "abcdabcdabed"
	s1 := "cd"

	// 调用标准库
	{
		index := strings.Index(s, s1)
		fmt.Println(index)
	}

	{
		index := simple(s, s1)
		fmt.Println(index)
	}
}

// simple
// 自己实现的简单 kmp，时间复杂度是 n*n1，暴力方法
// 简单的逻辑是遍历 s，发现字符和 s1的开头字符一致时，取之后整个长度的字符去判断是否一致
// TODO 只有核心功能实现，还可以优化判断长度，缩小一些时间
func simple(s, s1 string) int {
	n := len(s)
	n1 := len(s1)
	if n == 0 || n1 == 0 {
		return -1
	}
	for i, b := range s {
		if b == int32(s1[0]) {
			if s[i:i+n1] == s1 {
				return i
			}
		}
	}
	return -1
}
