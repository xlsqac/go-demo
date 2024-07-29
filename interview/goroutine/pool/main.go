// package main
// 实现一个协程池
package main

import (
	"errors"
	"log"
	"sync"
	"sync/atomic"
	"time"
)

// 协程池运行状态
// RUNNING 0 运行中
// STOP 1 已关闭，关闭的协程池不允许添加任务
const (
	RUNNING uint64 = iota
	STOP
)

// Task 执行的任务，就是一个函数的封装
type Task struct {
	Params  []any
	Handler func(params ...any)
}

// GoroutinePools 协程池
// capacity 当前池的容量
type GoroutinePools struct {
	capacity  uint64     // 容量
	workerNum uint64     // 运行中的任务数量
	status    uint64     // 池子的状态
	taskChan  chan *Task // 存放任务的通道

	sync.Mutex // 内嵌一个锁，保证添加任务时的同步安全
}

// getPoolCapacity 获取容量
func (p *GoroutinePools) getPoolCapacity() uint64 {
	return p.capacity
}

// getWorkerNum 获取任务数
func (p *GoroutinePools) getWorkerNum() uint64 {
	return atomic.LoadUint64(&p.workerNum)
}

// addNewWorker 增加一个任务数
func (p *GoroutinePools) addNewWorker() {
	atomic.AddUint64(&p.workerNum, 1)
}

// decWorker 减少一个任务数
func (p *GoroutinePools) decWorker() {
	atomic.AddUint64(&p.workerNum, -1)
}

// getPoolStatus 获取池子状态
func (p *GoroutinePools) getPoolStatus() uint64 {
	return atomic.LoadUint64(&p.status)
}

// setPoolStatus 设置池子状态，需要加锁操作
func (p *GoroutinePools) setPoolStatus(status uint64) bool {
	p.Lock()
	defer p.Unlock()

	if p.status == status {
		return false
	}
	p.status = status
	return true
}

// newTaskGoroutine 新增一个任务，从通道里面取任务出来，通过 recover 捕获任务的 panic
func (p *GoroutinePools) newTaskGoroutine() {
	p.addNewWorker()
	go func() {
		defer func() {
			p.decWorker()
			if r := recover(); r != nil {
				log.Printf("worker %s is panic\n", r)
			}
		}()

		for {
			select {
			case task, ok := <-p.taskChan:
				if !ok {
					return
				}
				task.Handler(task.Params...)
			}
		}
	}()
}

// newTask 新建一个任务，需要加锁，池子的状态得是 RUNNING
// 如果当前的任务数小于池子容量，可以添加任务
func (p *GoroutinePools) newTask(t *Task) error {
	p.Lock()
	defer p.Unlock()

	if p.getPoolStatus() == RUNNING {
		if p.getPoolCapacity() > p.getWorkerNum() {
			p.newTaskGoroutine()
		}
		p.taskChan <- t
		return nil
	} else {
		return errors.New("goroutine pool has already closed")
	}

}

// closePool 关闭协程池，改变池子状态并且通道中的任务数为 0 时关闭通道。
func (p *GoroutinePools) closePool() {
	p.setPoolStatus(STOP)
	for len(p.taskChan) > 0 {
		time.Sleep(time.Second * 60)
	}
	close(p.taskChan)
}

// main 测试一下协程池是否能正常工作 TODO
func main() {}
