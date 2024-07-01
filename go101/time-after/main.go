// package main
// 验证 time.After 导致的内存泄露问题
package main

import (
    "fmt"
    _ "net/http/pprof"
    "time"
)

func main() {
    //timer := time.NewTimer(time.Second * 5)
    //timer1 := time.NewTicker(time.Second * 3)
    //defer timer1.Stop()
    //time.After(time.Second)

    ch := make(chan int)
    go func() {
        in := 1
        for {
            in++
            ch <- in
        }
    }()

    //go func() {
    //	for {
    //timer.Reset(time.Second * 5)
    //select {
    //case _ = <-ch:
    //	continue
    //case <-timer1.C: // 只会触发一次 没有 reset 的话
    //	fmt.Println("time out", time.Now().Unix())
    //}
    //}
    //}()

    // After 会导致内存泄露
    //for {
    //	select {
    //	case _ = <-ch:
    //		continue
    //	case <-time.After(time.Second * 2):
    //		fmt.Println("time out after")
    //	}
    //}
    // 优化版
    ticker := time.NewTicker(time.Second * 2)
    defer ticker.Stop()

    for {
        select {
        case v := <-ch:
            fmt.Println(v)
            continue
        case <-ticker.C:
            fmt.Println("time out", time.Now().Unix())
        }
    }

    // timer AfterFunc
    //{
    //
    //	f := func() {
    //		fmt.Println("time after func")
    //	}
    //	t := time.AfterFunc(time.Second, f)
    //	defer t.Stop()
    //	select {
    //	case <-t.C:
    //		fmt.Println("time after channel")
    //	}
    //
    //}

    // ticker

    // NewTicker
    //{
    //	ticker := time.NewTicker(time.Second * 3)
    //	defer ticker.Stop()
    //	for {
    //		select {
    //		case <-ticker.C:
    //			fmt.Println("time out", time.Now().Unix())
    //		}
    //	}
    //}

    // Tick
    {
        tc := time.Tick(time.Second * 3)
        for {
            select {
            case <-tc:
                fmt.Println("time out", time.Now().Unix())
            }
        }
    }

    //_ = http.ListenAndServe("0.0.0.0:6060", nil)

}
