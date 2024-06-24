// package main
// https://leetcode.cn/problems/lru-cache/description/
package main

import "fmt"

type DLinkedListNode struct {
	Key   int
	Value int
	Prev  *DLinkedListNode
	Next  *DLinkedListNode
}

type LRUCache struct {
	cache    map[int]*DLinkedListNode
	capacity int
	size     int
	head     *DLinkedListNode
	tail     *DLinkedListNode
}

func initDLinkedListNode(key, value int) *DLinkedListNode {
	return &DLinkedListNode{
		Key:   key,
		Value: value,
	}
}

func Constructor(capacity int) *LRUCache {
	c := make(map[int]*DLinkedListNode)
	lRUCache := &LRUCache{
		cache:    c,
		capacity: capacity,
		head:     initDLinkedListNode(0, 0),
		tail:     initDLinkedListNode(0, 0),
	}
	// 真实的头尾是不会变的，每次操作的是 head.next 和 tail.prev
	lRUCache.head.Next = lRUCache.tail
	lRUCache.tail.Prev = lRUCache.head
	return lRUCache
}

func (l *LRUCache) Get(key int) int {
	if _, ok := l.cache[key]; !ok {
		return -1
	}
	node := l.cache[key]
	l.moveToHead(node)
	return node.Value
}

func (l *LRUCache) Put(key int, value int) {
	if _, ok := l.cache[key]; !ok {
		if l.size == l.capacity {
			node := l.removeTail()
			delete(l.cache, node.Key)
			l.size--
		}
		node := initDLinkedListNode(key, value)
		l.addToHead(node)
		l.cache[key] = node
		l.size++
		return
	}
	node := l.cache[key]
	node.Value = value
	l.moveToHead(node)
}

func (l *LRUCache) addToHead(node *DLinkedListNode) {
	node.Prev = l.head
	node.Next = l.head.Next
	l.head.Next.Prev = node
	l.head.Next = node
}

func (l *LRUCache) moveToHead(node *DLinkedListNode) {
	l.removeNode(node)
	l.addToHead(node)
}

func (l *LRUCache) removeNode(node *DLinkedListNode) {
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
}

func (l *LRUCache) removeTail() *DLinkedListNode {
	node := l.tail.Prev
	l.removeNode(node)
	return node
}

func main() {
	l := Constructor(4)
	l.Put(1, 1)
	l.Put(2, 2)
	l.Put(3, 3)
	l.Put(4, 4)
	fmt.Println(l.Get(1))
	fmt.Println(l.Get(2))
	fmt.Println(l.Get(3))
	fmt.Println(l.Get(4))
	l.Put(5, 5)
	fmt.Println(l.Get(5))
	fmt.Println(l.Get(1))
}
