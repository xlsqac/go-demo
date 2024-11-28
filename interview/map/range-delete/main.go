// package main
// 边遍历边删除元素
// 边遍历边删除会删除掉原 map 中的元素，不会删除当前遍历的 map，因为当前遍历的 map 是一个副本
package main

import "fmt"

func main() {
	myMap := map[string]int{
		"key1":  1,
		"key2":  2,
		"key3":  3,
		"key4":  4,
		"key5":  5,
		"key6":  6,
		"key7":  7,
		"key8":  8,
		"key9":  9,
		"key10": 10,
		"key11": 11,
		"key12": 12,
		"key13": 13,
		"key14": 14,
		"key15": 15,
		"key16": 16,
		"key17": 17,
		"key18": 18,
		"key19": 19,
		"key20": 20,
	}
	for k, v := range myMap {
		if v%2 == 0 {
			delete(myMap, k)
		}
		fmt.Println(k, v)
	}

	fmt.Println("-------------------")
	for k, v := range myMap {
		fmt.Println(k, v)
	}

}
