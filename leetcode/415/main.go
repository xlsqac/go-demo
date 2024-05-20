//	Package main
//
// 大数相加 https://leetcode.cn/problems/add-strings/description/
// 给定两个字符串形式的非负整数 num1 和num2 ，计算它们的和并同样以字符串形式返回
// 把两个字符串每位相加，如果长度不一致的话在短串前补零
// 通过取余和整除的方式来确定是否需要进位
// 采用只遍历短串，最后把长串未遍历到的部分和相加过的拼在一起的话，需要考虑拼接时的进位问题
// 拼接时进位需要递归实现，会拉低平均时间复杂度，所以一般不考虑采用此方法实现
// 时间复杂度 O(max(len(num1), len(num2)))
package main

import (
	"fmt"
	"strings"
)

func padZeros(num1, num2 string) (string, string) {
	lenA, lenB := len(num1), len(num2)
	maxLen := max(lenA, lenB)
	padding := maxLen - lenA
	padStr := "0"

	if padding > 0 {
		num1 = strings.Repeat(padStr, padding) + num1
	} else {
		num2 = strings.Repeat(padStr, maxLen-lenB) + num2
	}

	return num1, num2
}

func addStrings(num1 string, num2 string) string {
	// 将 num1，num2 的长度设为相等，给少的一方前面补 0
	num1, num2 = padZeros(num1, num2)
	length := len(num1)
	// 存放 ascii 码，length+1 是因为想加时可能会进位，进位最多进1位
	result := make([]byte, length+1)
	// 进位
	carry := byte(0)

	for i := length - 1; i >= 0; i-- {
		digitA := num1[i] - '0'
		digitB := num2[i] - '0'
		// carry = 1 说明上一位相加是需要进位
		sum := digitA + digitB + carry
		// sum % 10 保留个位
		result[i+1] = sum%10 + '0'
		// 进位值，下次循环会把这个值加到和中
		carry = sum / 10
	}
	// 加上最高位相加时的进位
	result[0] = carry + '0'
	// 如果是 0 说明最终的结果没有导致长度进位，所以不需要这个
	if result[0] == '0' {
		return string(result[1:])
	}
	return string(result)
}

func main() {
	num1 := "11"
	num2 := "123"
	sum := addStrings(num1, num2)
	// sum == 533
	fmt.Println(sum)
}
