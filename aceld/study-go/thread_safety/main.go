package main

import (
	"fmt"
	"sync"
	"time"
)

var sum int
var sum1 int

var sum2 int
var wait sync.WaitGroup

var wait1 sync.WaitGroup
var lock sync.Mutex

func add() {
	for i := 0; i < 1000000; i++ {
		sum++
	}
}

func sub() {
	for i := 0; i < 1000000; i++ {
		sum--
	}
}

func add1() {
	for i := 0; i < 1000000; i++ {
		sum1++
	}
	wait.Done()
}

func sub1() {
	for i := 0; i < 1000000; i++ {
		sum1--
	}
	wait.Done()
}

func add2() {
	lock.Lock()
	for i := 0; i < 1000000; i++ {
		sum2++
	}
	lock.Unlock()
	wait1.Done()
}

func sub2() {
	lock.Lock()
	for i := 0; i < 1000000; i++ {
		sum2--
	}
	lock.Unlock()
	wait1.Done()
}

func main() {
	startTime := time.Now()
	add()
	sub()
	fmt.Println("sync sum", sum, "time", time.Since(startTime))

	startTime1 := time.Now()
	wait.Add(2)
	go add1()
	go sub1()
	wait.Wait()
	fmt.Println("asynchronous not lock sum1", sum1, "time1", time.Since(startTime1))

	startTime2 := time.Now()
	wait1.Add(2)
	go add2()
	go sub2()
	wait1.Wait()
	fmt.Println("asynchronous lock sum2", sum2, "time2", time.Since(startTime2))

}
