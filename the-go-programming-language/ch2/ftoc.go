// 华氏度转换摄氏度，转换算法用函数封装
package main

import "fmt"

func main() {
	const freezingF, boilingF = 32.0, 212.0
	fmt.Printf("%g°F = %g°C\n", freezingF, fTOc(freezingF))
	fmt.Printf("%g°F = %g°C\n", boilingF, fTOc(boilingF))
}

func fTOc(f float64) float64 {
	return (f - 32) * 5 / 9
}
