// package main
package main

import (
	"fmt"
	"sync"
	"time"
)

type SliceQueue struct {
	data []interface{}
	mu   sync.Mutex
}

func NewSliceQueue(size int) *SliceQueue {
	return &SliceQueue{data: make([]interface{}, 0, size)}
}

func (s *SliceQueue) Enqueue(v interface{}) {
	s.mu.Lock()
	s.data = append(s.data, v)
	s.mu.Unlock()
}

func (s *SliceQueue) Dequeue() interface{} {
	s.mu.Lock()
	if len(s.data) == 0 {
		s.mu.Unlock()
		return nil
	}
	v := s.data[0]
	s.data = s.data[1:]
	s.mu.Unlock()
	return v
}

func main() {
	sliceQueue := NewSliceQueue(10)
	for i := 0; i < 1000; i++ {
		go func() {
			sliceQueue.Enqueue(i)
		}()
	}
	for i := 0; i < 1000; i++ {
		go func() {
			sliceQueue.Dequeue()
		}()
	}
	time.Sleep(time.Second * 1)
	fmt.Println(sliceQueue.data)

}
