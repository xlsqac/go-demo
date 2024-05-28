package main

import "fmt"

func findWinners(matchs [][]int) [][]int {
	var answer = [][]int{}
	// 每个人赢的次数
	var personWinTimes = make(map[int]int)
	for _, v := range matchs {
		personWinTimes[v[0]]++
		personWinTimes[v[1]]--
	}
	fmt.Println(personWinTimes)
	return answer

}

func main() {
	matches := [][]int{
		{2, 3},
		{1, 3},
		{5, 4},
		{6, 4},
	}
	answer := findWinners(matches)
	fmt.Println(answer)

}
