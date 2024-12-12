// package main
// https://leetcode.cn/problems/longest-substring-without-repeating-characters/?envType=study-plan-v2&envId=top-100-liked
// 无重复字符的最长子串
package main

import "fmt"

func main() {
	{
		s := "abcabcbb"
		fmt.Println(3, lengthOfLongestSubstring2(s))
		// a b 0 1 0 2 2-0=2 ab
		// a c 0 2 0 3 3-0=3 abc
		// a a 0 3 1 3
		// b a 1 3 1 4 4-1=3 bca
		// b b 1 4 2 4 4-2=2 bca
		// c b 2 4 2 5 5-2=3 cab
		// c c 2 5 3 5
		// a c 3 5 3 6 6-3=3 abc
		// a b 3 6 3 7 7-3=4 abcb
	}
	{
		s := "bbbbb"
		fmt.Println(1, lengthOfLongestSubstring2(s))
	}
	{
		s := "pwwkew"
		fmt.Println(3, lengthOfLongestSubstring2(s))
	}
	{
		s := "ab"
		fmt.Println(2, lengthOfLongestSubstring2(s))
	}
}

// lengthOfLongestSubstring
func lengthOfLongestSubstring(s string) int {
	lenS := len(s)
	if lenS < 2 {
		return lenS
	}
	ans := 1
	left := 0
	right := 1
	for left < right && right < lenS {
		strMap := make(map[byte]struct{})
		for i := left; i < right; i++ {
			// fmt.Println(left, right)
			if _, ok := strMap[s[i]]; ok {
				// fmt.Println(left, right)
				left++
				break
			}
			strMap[s[i]] = struct{}{}
		}

		if _, ok := strMap[s[right]]; ok {
			// fmt.Println(left, right)
			left++
			right++
		} else {
			strMap[s[right]] = struct{}{}
			right++
			// fmt.Println(left, right)
			ans = max(ans, right-left)
		}
	}
	return ans
}

// lengthOfLongestSubstring2
// left 指针每次重复后往右移动
// right 指针每次循环后往右移动
func lengthOfLongestSubstring2(s string) int {
	ans := 0
	counter := make(map[rune]int)
	left := 0
	for right, char := range s {
		counter[char]++
		for counter[char] > 1 {
			counter[rune(s[left])]--
			left++
		}
		ans = max(ans, right-left+1)
	}
	return ans
}
