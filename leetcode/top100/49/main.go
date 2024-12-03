// package main
// https://leetcode.cn/problems/group-anagrams/description/?envType=study-plan-v2&envId=top-100-liked
// 字母异位词分组
package main

import (
	"fmt"
	"slices"
)

func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Println(groupAnagrams(strs))

	strs = []string{""}
	fmt.Println(groupAnagrams(strs))

	strs = []string{"a"}
	fmt.Println(groupAnagrams(strs))

	strs = []string{"cab", "tin", "pew", "duh", "may", "ill", "buy", "bar", "max", "doc"}
	fmt.Println(groupAnagrams(strs))
}

// groupAnagrams return [][]string 所有的字母异位词
// 先遍历每个str，然后给 str 排序，排完序直接作为 key
// 然后把异位词放到一个 key 里
// 最后再遍历 map 得到结果
func groupAnagrams(strs []string) [][]string {
	sortedStrMap := map[string][]string{}

	for _, str := range strs {
		s := []byte(str)
		slices.Sort(s)
		sortedStr := string(s)
		sortedStrMap[sortedStr] = append(sortedStrMap[sortedStr], str)
	}

	res := make([][]string, 0, len(sortedStrMap))
	for _, v := range sortedStrMap {
		res = append(res, v)
	}
	return res
}
