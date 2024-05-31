// package
// https://leetcode.cn/problems/find-longest-special-substring-that-occurs-thrice-i/description/
// TODO
package main

import "fmt"

func maximumLength(s string) int {
	// gen map
	m := map[string][]int{}
	for index, str := range s {
		k := string(str)
		m[k] = append(m[k], index)
	}
	fmt.Println(m)
	// map[a:[0 3 5] b:[1 4] c:[2]]
	// map[a:[0 1 2 3]]
	// map[a:[0] b:[1] c:[2 3 4 5 6] d:[7 8 9 10]]
	ans, times := -1, 3
	for _, vec := range m {
		left, right := 0, 0
		// 出现小于三次
		if len(vec) < times {
			continue
		} else if len(vec) == times {
			ans = 1
		}
		// 理想结果的长度，每个字节都挨着
		//perfectAns := len(vec) - 2
		for right = 0; right < len(vec); right++ {
			// 只要出现间隔大于 1，就说明没有挨着, 长度就为 1
			if vec[right]-vec[left] > 1 {
				left++
			}
			// 只要存在 3 次，最小值为 1
			ans = max(ans, right-left+1)
		}
	}

	return ans
}

func main() {
	s := "aaaa"
	maximum := maximumLength(s)
	fmt.Println(maximum, maximum == 2)

	s = "abcdef"
	maximum = maximumLength(s)
	fmt.Println(maximum, maximum == -1)

	s = "abcaba"
	maximum = maximumLength(s)
	fmt.Println(maximum, maximum == 1)

	s = "abcccccdddd"
	maximum = maximumLength(s)
	fmt.Println(maximum, maximum == 3)
}
