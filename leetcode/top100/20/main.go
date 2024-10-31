// package main
// https://leetcode.cn/problems/valid-parentheses/description/?envType=study-plan-v2&envId=top-100-liked
// 有效的括号
package main

import "fmt"

func main() {
	{
		s := "()"
		fmt.Println(isValid(s))
	}
	{
		s := "()[]{}"
		fmt.Println(isValid(s))
	}
	{
		s := "(]"
		fmt.Println(isValid(s))
	}
	{
		s := "([])"
		fmt.Println(isValid(s))
	}

	{
		s := "([)]"
		fmt.Println(isValid(s))
	}
}

func isValid(s string) bool {
	if len(s)%2 != 0 {
		return false
	}

	pairs := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}
	stack := []byte{}

	for i := 0; i < len(s); i++ {
		// 如果是右括号，判断栈顶的元素是否是对应的左括号
		// 如果是左括号入栈
		if pairs[s[i]] > 0 {
			// 如果栈是空的，或者栈顶的元素不是对应的左括号，返回 false
			if len(stack) == 0 || stack[len(stack)-1] != pairs[s[i]] {
				return false
			}
			// 如果有对应的左括号，那么将其出栈
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, s[i])
		}
	}
	return len(stack) == 0
}
