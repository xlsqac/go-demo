package main

import (
    "fmt"
    "sync"
    "time"
)

var moneyChan = make(chan int)
var nameChan = make(chan string)
var doneChan = make(chan int)

func sopping(name string, money int, wait *sync.WaitGroup) {
    fmt.Println(name, "sopping")
    time.Sleep(1 * time.Second)
    moneyChan <- money
    nameChan <- name
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
        defer close(doneChan)
        wait.Wait()
    }()

    var moneyList []int
    var nameList []string

    //for {
    select {
    case money := <-moneyChan:
        moneyList = append(moneyList, money)
    case name := <-nameChan:
        nameList = append(nameList, name)
    case <-doneChan:
        fmt.Println("moneyList", moneyList)
        fmt.Println("nameList", nameList)
        fmt.Println("end", time.Since(startTime))
        return
    }
    //}
    chan1 := make(chan int)
    chan2 := make(chan int)

    go func() {
        chan1 <- 1
    }()

    go func() {
        chan2 <- 1
    }()

    select {
    case <-chan1:
        fmt.Println("chan1 ready.")
    case <-chan2:
        fmt.Println("chan2 ready.")
    }

}
