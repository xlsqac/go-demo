package main

/*
数组
1. 随机访问
*/
import (
	"fmt"
	"math/rand"
)

func randomAccess(arr [5]int) int {
	// 随机访问，O(1)
	index := rand.Intn(len(arr))
	return arr[index]

}

func insert(arr [5]int, index int, value int) [5]int {
	// 插入元素，O(n)
	for i := len(arr) - 1; i > index; i-- {
		arr[i] = arr[i-1]
	}
	arr[index] = value
	return arr

}

func remove(arr [5]int, index int) [5]int {
	// 删除元素, O(n)
	for i := index; i < len(arr)-1; i++ {
		arr[i] = arr[i+1]
	}
	return arr
}

func traverse(arr [5]int) {
	// 遍历数组, O(n)
	for i := 0; i < len(arr); i++ {
		fmt.Printf("%d ", arr[i])
	}
}

func find(arr [5]int, value int) int {
	//查找元素，找到元素返回索引，不存在返回 -1. O(n)
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] == value {
			return i
		}
	}
	return -1
}

func extend(arr []int, enlarge int) []int {
	//扩容，数组不方便实现扩容，这里用 slice
	res := make([]int, len(arr)+enlarge)
	for i, num := range arr {
		res[i] = num
	}
	return res
}

func main() {
	var arr = [5]int{1, 3, 2, 5, 4}
	fmt.Printf("[array] %v\n", arr)

	randomNum := randomAccess(arr)
	fmt.Printf("[random access] %d\n", randomNum)

	insArr := insert(arr, 2, 10)
	fmt.Printf("[inserted array] %v\n", insArr)

	remArr := remove(arr, 2)
	fmt.Printf("[removed array] %v\n", remArr)

	fmt.Print("[traverse array] ")
	traverse(arr)
	fmt.Print("\n")

	hasIndex := find(arr, 3)
	fmt.Printf("[find 3 output 1] %d\n", hasIndex)

	notHasIndex := find(arr, 100)
	fmt.Printf("[find 100 output -1] %d\n", notHasIndex)

	extendArr := []int{1, 3, 2, 5, 4}
	enlarge := 10
	newArr := extend(extendArr, enlarge)
	fmt.Printf("[extend array] %v\n", newArr)
	fmt.Printf("[extend array length] %d\n", len(extendArr)+enlarge)

}
