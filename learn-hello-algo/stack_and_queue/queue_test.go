package stack_and_queue

import (
	"container/list"
	"fmt"
	"testing"
)

//func PrintList(list *list.List) {
//	if list.Len() == 0 {
//		fmt.Print("[]\n")
//		return
//	}
//	e := list.Front()
//	// 强转为 string, 会影响效率
//	fmt.Print("[")
//	for e.Next() != nil {
//		fmt.Print(e.Value, " ")
//		e = e.Next()
//	}
//	fmt.Print(e.Value, "]\n")
//}

func PrintSlice[T any](nums []T) {
	fmt.Printf("%v", nums)
	fmt.Println()
}

/* 直接将 list 做为队列使用 */
func TestListQueue(t *testing.T) {
	queue := list.New()

	/* 元素入队 */
	queue.PushBack(1)
	queue.PushBack(3)
	queue.PushBack(2)
	queue.PushBack(5)
	queue.PushBack(4)
	fmt.Print("[list queue to list] ")
	PrintList(queue)

	/* 访问队首元素 */
	peek := queue.Front()
	fmt.Printf("[list queue peek] %d\n", peek.Value)

	/* 出队 */
	pop := queue.Front()
	queue.Remove(pop)
	fmt.Printf("[list queue pop] %d\n", pop.Value)

	/* 获取队列长度 */
	size := queue.Len()
	fmt.Printf("[list queue size] %d\n", size)

	/* 判断队列是否为空 */
	isEmpty := size == 0
	fmt.Printf("[list queue is empty] %t\n", isEmpty)
}

func TestArrayQueue(t *testing.T) {
	// 初始化队列，使用队列的通用接口
	capacity := 10
	queue := newArrayQueue(capacity)

	// 元素入队
	queue.push(1)
	queue.push(3)
	queue.push(2)
	queue.push(5)
	queue.push(4)
	fmt.Print("队列 queue = ")
	PrintSlice(queue.toSlice())

	// 访问队首元素
	peek := queue.peek()
	fmt.Println("队首元素 peek =", peek)

	// 元素出队
	pop := queue.pop()
	fmt.Print("出队元素 pop = ", pop, ", 出队后 queue = ")
	PrintSlice(queue.toSlice())

	// 获取队的长度
	size := queue.size()
	fmt.Println("队的长度 size =", size)

	// 判断是否为空
	isEmpty := queue.isEmpty()
	fmt.Println("队是否为空 =", isEmpty)

	/* 测试环形数组 */
	for i := 0; i < 10; i++ {
		queue.push(i)
		queue.pop()
		fmt.Print("第", i, "轮入队 + 出队后 queue =")
		PrintSlice(queue.toSlice())
	}
}

func TestLinkedListQueue(t *testing.T) {
	// 初始化队
	queue := newLinkedListQueue()

	// 元素入队
	queue.push(1)
	queue.push(3)
	queue.push(2)
	queue.push(5)
	queue.push(4)
	fmt.Print("队列 queue = ")
	PrintList(queue.toList())

	// 访问队首元素
	peek := queue.peek()
	fmt.Println("队首元素 peek =", peek)

	// 元素出队
	pop := queue.pop()
	fmt.Print("出队元素 pop = ", pop, ", 出队后 queue = ")
	PrintList(queue.toList())

	// 获取队的长度
	size := queue.size()
	fmt.Println("队的长度 size =", size)

	// 判断是否为空
	isEmpty := queue.isEmpty()
	fmt.Println("队是否为空 =", isEmpty)
}

// BenchmarkArrayQueue 8 ns/op in Mac M1 Pro
func BenchmarkArrayQueue(b *testing.B) {
	capacity := 1000
	queue := newArrayQueue(capacity)
	// use b.N for looping
	for i := 0; i < b.N; i++ {
		queue.push(777)
	}
	for i := 0; i < b.N; i++ {
		queue.pop()
	}
}

// BenchmarkLinkedQueue 62.66 ns/op in Mac M1 Pro
func BenchmarkLinkedQueue(b *testing.B) {
	queue := newLinkedListQueue()
	// use b.N for looping
	for i := 0; i < b.N; i++ {
		queue.push(777)
	}
	for i := 0; i < b.N; i++ {
		queue.pop()
	}
}
