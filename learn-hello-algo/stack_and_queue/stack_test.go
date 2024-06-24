package stack_and_queue

import (
	"container/list"
	"fmt"
	"testing"
)

func PrintList(list *list.List) {
	if list.Len() == 0 {
		fmt.Print("[]\n")
		return
	}
	e := list.Front()
	// 强转为 string, 会影响效率
	fmt.Print("[")
	for e.Next() != nil {
		fmt.Print(e.Value, " ")
		e = e.Next()
	}
	fmt.Print(e.Value, "]\n")
}

func TestStack(t *testing.T) {
	// 直接用 slice 做栈
	var stack []int
	// 入栈
	stack = append(stack, 1)
	stack = append(stack, 3)
	stack = append(stack, 2)
	stack = append(stack, 5)
	stack = append(stack, 4)
	t.Logf("[sliceStack] %v", stack)

	// 获取栈的长度
	size := len(stack)
	fmt.Println("栈的长度 size =", size)

	// 访问栈顶元素
	peek := stack[size-1]
	fmt.Println("栈顶元素 peek =", peek)

	// 出栈
	pop := stack[size-1]
	stack = stack[:size-1]
	fmt.Print("出栈元素 pop = ", pop, ", 出栈后 stack = ")
	t.Logf("[sliceStack] %v", stack)

	// 判断是否为空
	isEmpty := len(stack) == 0
	fmt.Println("栈是否为空 =", isEmpty)

}

func TestArrayStack(t *testing.T) {
	arrayStack := NewArrayStack()
	arrayStack.Push(1)
	arrayStack.Push(3)
	arrayStack.Push(2)
	arrayStack.Push(5)
	arrayStack.Push(4)

	t.Logf("[arrayStack] %v", arrayStack.ToSlice())

	// 访问栈顶元素
	peek := arrayStack.Peek()
	fmt.Println("栈顶元素 peek =", peek)

	// 元素出栈
	pop := arrayStack.Pop()
	fmt.Print("出栈元素 pop = ", pop, ", 出栈后 stack = ")
	t.Logf("[arrayStack] %v", arrayStack.ToSlice())

	// 获取栈的长度
	size := arrayStack.Size()
	fmt.Println("栈的长度 size =", size)

	// 判断是否为空
	isEmpty := arrayStack.IsEmpty()
	fmt.Println("栈是否为空 =", isEmpty)
}

func TestLinkedListStack(t *testing.T) {
	linkedListStack := NewLinkedListStack()
	linkedListStack.Push(1)
	linkedListStack.Push(3)
	linkedListStack.Push(2)
	linkedListStack.Push(5)
	linkedListStack.Push(4)

	t.Log("[LinkedListStack]")
	PrintList(linkedListStack.ToList())

	// 访问栈顶元素
	peek := linkedListStack.Peek()
	fmt.Println("栈顶元素 peek =", peek)

	// 元素出栈
	pop := linkedListStack.Pop()
	fmt.Print("出栈元素 pop = ", pop, ", 出栈后 stack = ")
	PrintList(linkedListStack.ToList())

	// 获取栈的长度
	size := linkedListStack.Size()
	fmt.Println("栈的长度 size =", size)

	// 判断是否为空
	isEmpty := linkedListStack.IsEmpty()
	fmt.Println("栈是否为空 =", isEmpty)
}

func BenchmarkArrayStack(b *testing.B) {
	stack := NewArrayStack()
	// use b.N for looping
	for i := 0; i < b.N; i++ {
		stack.Push(777)
	}
	for i := 0; i < b.N; i++ {
		stack.Pop()
	}
}

func BenchmarkLinkedListStack(b *testing.B) {
	stack := NewLinkedListStack()
	// use b.N for looping
	for i := 0; i < b.N; i++ {
		stack.Push(777)
	}
	for i := 0; i < b.N; i++ {
		stack.Pop()
	}
}
