package main

import (
	"fmt"
	"sync"
	"time"
)

var moneyChan = make(chan int)
var nameChan = make(chan string)

func sopping(name string, money int, wait *sync.WaitGroup) {
	fmt.Println(name, "sopping")
	time.Sleep(1 * time.Second)
	moneyChan <- money
	fmt.Println("save money", money)
	nameChan <- name
	fmt.Println("save name", name)
	wait.Done()
}

func main() {

	var wait sync.WaitGroup
	startTime := time.Now()

	wait.Add(3)
	go sopping("1", 1, &wait)
	go sopping("2", 2, &wait)
	go sopping("3", 3, &wait)

	go func() {
		defer close(moneyChan)
		defer close(nameChan)
		wait.Wait()
	}()

	var moneyList []int
	var nameList []string

	//go func() {
	//    for money := range moneyChan {
	//        moneyList = append(moneyList, money)
	//    }
	//}()
	fmt.Println("start get chan data", time.Since(startTime))

	// 死锁的原因是 chan 没有长度，所以必须传一个取一个
	// 但是 shopping 中会给两个 chan 传数据，而 nameChan 必须等 moneyChan 取完只有才开始取
	// 所以就会阻塞在 nameChan <- name，所有协程都没办法 done，chan 也就没办法 close，导致 range moneyChan 一直在等待数据
	// 可以改成一个异步取，一个阻塞的取或者给 chan 一个大于等于协程个数的长度
	for money := range moneyChan {
		fmt.Println("money", money)
		moneyList = append(moneyList, money)
	}
	fmt.Println("moneyList", moneyList)

	for name := range nameChan {
		nameList = append(nameList, name)
	}
	fmt.Println("nameList", nameList)
	fmt.Println("end", time.Since(startTime))

}
