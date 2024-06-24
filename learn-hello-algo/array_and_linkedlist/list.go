// 列表及其相关操作，基于 slice 实现
package main

import (
    "fmt"
    "sort"
)

func main() {
    // 未初始化和初始化的列表
    // nums1 := []int{}
    nums := []int{1, 3, 2, 5, 4}

    // 访问元素 access O(1)
    fmt.Printf("[access 2] %d\n", nums[2])

    // 插入元素 insert O(n)
    // 索引 2 插入数字 10
    nums = append(nums[:2], append([]int{6}, nums[2:]...)...)
    fmt.Printf("[append index: 2 value: 6] %v\n", nums)

    // 追加元素 append O(1) 10
    nums = append(nums, 10)
    fmt.Printf("[append -1] %d\n", nums[6])

    // 根据索引删除元素 O(n) 2
    nums = append(nums[:2], nums[3:]...)
    fmt.Printf("[pop] %v\n", nums)

    // 遍历列表 reverse O(n)
    fmt.Print("[reverse]")
    for _, num := range nums {
        fmt.Printf("%d ", num)
    }
    fmt.Print("\n")

    // 拼接元素 extend O(n)
    nums1 := []int{6, 7}
    nums = append(nums, nums1...)
    fmt.Printf("[extend] %v\n", nums)

    // 排序
    sort.Ints(nums)
    fmt.Printf("[sort] %v\n", nums)

}
