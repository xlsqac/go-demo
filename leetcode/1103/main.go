// package main
// https://leetcode.cn/problems/distribute-candies-to-people/description/
package main

import (
	"fmt"
)

func distributeCandies(candies int, numPeople int) []int {
	ans := make([]int, numPeople)
	for i := 0; candies > 0; i++ {
		ans[i%numPeople] += min(candies, i+1)
		candies -= min(candies, i+1)

	}
	return ans

}

func main() {
	candies, numPeople := 7, 4
	res := distributeCandies(candies, numPeople)
	fmt.Println(res)
	fmt.Println()

	candies, numPeople = 10, 3
	res = distributeCandies(candies, numPeople)
	fmt.Println(res)
}
