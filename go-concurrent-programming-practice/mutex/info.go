// package main
// 根据 Mutex.state 的值，获取锁的状态
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
	"unsafe"
)

const (
	mutexLocked = 1 << iota // mutex is locked
	mutexWoken
	mutexStarving
	mutexWaiterShift = iota
)

type Mutex struct {
	sync.Mutex
}

func (m *Mutex) count() int {
	v := atomic.LoadInt32((*int32)(unsafe.Pointer(&m.Mutex))) // 获取 state 的值
	v = v >> mutexWaiterShift                                 // 获取等待者数量的值
	v = v + (v & mutexLocked)                                 // 加上锁持有者的数量
	return int(v)
}

func (m *Mutex) IsLocked() bool {
	v := atomic.LoadInt32((*int32)(unsafe.Pointer(&m.Mutex))) // 获取 state 的值
	return v&mutexLocked == mutexLocked
}

func (m *Mutex) IsWoken() bool {
	v := atomic.LoadInt32((*int32)(unsafe.Pointer(&m.Mutex))) // 获取 state 的值
	return v&mutexWoken == mutexWoken
}

func (m *Mutex) IsStarving() bool {
	v := atomic.LoadInt32((*int32)(unsafe.Pointer(&m.Mutex))) // 获取 state 的值
	return v&mutexStarving == mutexStarving
}

func main() {
	var mu Mutex
	for i := 0; i < 1000; i++ {
		go func() {
			mu.Lock()
			time.Sleep(time.Second * 1)
			mu.Unlock()
		}()
	}

	time.Sleep(time.Second * 1)
	fmt.Printf("waitings: %d, isLocked: %t, isWoken: %t, isStarving: %t\n", mu.count(), mu.IsLocked(), mu.IsWoken(), mu.IsStarving())
}
