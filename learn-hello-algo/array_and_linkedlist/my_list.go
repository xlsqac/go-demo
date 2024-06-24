/*
基于 slice 实现一个列表
*/
package main

import "fmt"

/* 列表结构体 */
type myList struct {
    arrCapacity int
    arr         []int
    arrSize     int
    extendRatio int
}

/* 构造函数 */
func newMyList() *myList {
    return &myList{
        arrCapacity: 10,              // 列表容量
        arr:         make([]int, 10), // 数组
        arrSize:     0,               // 列表长度
        extendRatio: 2,               // 扩容倍数
    }
}

/* 获取长度 */
func (l *myList) size() int {
    return l.arrSize
}

/* 获取容量 */
func (l *myList) capacity() int {
    return l.arrCapacity
}

/* 访问元素 */
func (l *myList) get(index int) int {
    l.checkIndex(index)
    return l.arr[index]

}

/* 更新元素 */
func (l *myList) set(index int, num int) {
    l.checkIndex(index)
    l.arr[index] = num
}

/* 在尾部添加元素 */
func (l *myList) add(num int) {
    l.checkIsNeedExtendCapacity()
    l.arr[l.arrSize] = num
    l.arrSize++
}

/* 在中间插入元素 */
func (l *myList) insert(index int, num int) {
    l.checkIndex(index)
    l.checkIsNeedExtendCapacity()
    for i := l.size() - 1; i >= index; i-- {
        l.arr[i+1] = l.arr[i]
    }
    l.arr[index] = num
    l.arrSize++

}

/* 删除元素 */
func (l *myList) remove(index int) int {
    l.checkIndex(index)
    num := l.arr[index]
    for i := index; i < l.size()-1; i++ {
        l.arr[i] = l.arr[i+1]
    }
    l.arrSize--
    return num
}

/* 返回有效长度的列表 */
func (l *myList) toArray() []int {
    return l.arr[:l.size()]
}

/* 检查是否需要扩容 */
func (l *myList) checkIsNeedExtendCapacity() {
    if l.size() == l.capacity() {
        l.extendCapacity()
    }
}

/* 扩容 */
func (l *myList) extendCapacity() {
    l.arr = append(l.arr, make([]int, l.arrCapacity*(l.extendRatio-1))...)
    l.arrCapacity = len(l.arr)
}

/* 检查索引 */
func (l *myList) checkIndex(index int) {
    if index < 0 || index >= l.size() {
        panic("索引越界")
    }
}

func main() {
    nums := newMyList()
    /* 在尾部添加元素 */
    nums.add(1)
    nums.add(3)
    nums.add(2)
    nums.add(5)
    nums.add(4)
    fmt.Printf("列表 nums = %v ，容量 = %v ，长度 = %v\n", nums.toArray(), nums.capacity(), nums.size())

    /* 在中间插入元素 */
    nums.insert(3, 6)
    fmt.Printf("在索引 3 处插入数字 6 ，得到 nums = %v\n", nums.toArray())

    /* 删除元素 */
    nums.remove(3)
    fmt.Printf("删除索引 3 处的元素，得到 nums = %v\n", nums.toArray())

    /* 访问元素 */
    num := nums.get(1)
    fmt.Printf("访问索引 1 处的元素，得到 num = %v\n", num)

    /* 更新元素 */
    nums.set(1, 0)
    fmt.Printf("将索引 1 处的元素更新为 0 ，得到 nums = %v\n", nums.toArray())

    /* 测试扩容机制 */
    for i := 0; i < 10; i++ {
        // 在 i = 5 时，列表长度将超出列表容量，此时触发扩容机制
        nums.add(i)
    }
    fmt.Printf("扩容后的列表 nums = %v ，容量 = %v ，长度 = %v\n", nums.toArray(), nums.capacity(), nums.size())
}
