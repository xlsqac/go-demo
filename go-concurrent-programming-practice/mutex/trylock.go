// package main
// 给 Mutex 增加 TryLock 方法
package main

import (
	"fmt"
	"math/rand"
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

func (m *Mutex) TryLock() bool {
	// 如果能成功抢到锁
	if atomic.CompareAndSwapInt32((*int32)(unsafe.Pointer(&m.Mutex)), 0, mutexLocked) {
		return true
	}

	// 如果锁被其他 g 持有或者准备持有就返回 false 不准备获取
	old := atomic.LoadInt32((*int32)(unsafe.Pointer(&m.Mutex)))
	if old&(mutexLocked|mutexStarving|mutexWoken) != 0 {
		return false
	}

	// 可以尝试获取锁
	new_ := old | mutexLocked
	return atomic.CompareAndSwapInt32((*int32)(unsafe.Pointer(&m.Mutex)), old, new_)
}

func main() {
	var mu Mutex

	go func() {
		mu.Lock()
		time.Sleep(time.Duration(rand.Intn(2)) * time.Second)
		//time.Sleep(time.Second * 3)
		mu.Unlock()
	}()
	time.Sleep(time.Second)

	ok := mu.TryLock()
	if ok {
		fmt.Println("get the lock")
		mu.Unlock()
		return
	}
	fmt.Println("can't get the lock")

}
