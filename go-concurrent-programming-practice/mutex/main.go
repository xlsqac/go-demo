// package main
// 启动 10 个 goroutine 对一个变量进行累加，每个 g 加 100000 次，不加锁的话会出现计数错误，运行时间大概是 800微秒到 1毫秒
// 加上锁是 60ms 左右
// 如果不使用锁，仅仅针对该问题，设置 CPU 数目为 1也可以解决并且时间大概在 1.7ms 左右
package main

import (
	"fmt"
	"sync"
	"time"
)

type Counter struct {
	sync.Mutex
	Count int64
}

func main() {
	// 骚操作，直接给 1 个核解决这个问题，项目中不能这样用
	//runtime.GOMAXPROCS(1)
	/*
		{
			start := time.Now()
			var mu sync.Mutex
			var count = 0
			var wg sync.WaitGroup
			wg.Add(10)
			for i := 0; i < 10; i++ {
				go func() {
					defer wg.Done()
					for j := 0; j < 100000; j++ {
						mu.Lock()
						count++ // 非原子操作，好多计数会被吞掉
						mu.Unlock()
					}
				}()
			}
			wg.Wait()
			fmt.Println(time.Since(start))
			fmt.Println(count) // 最终的结果可能不是1000000，因为有些goroutine的计数被吞掉了，实测 300000 左右
		}
	*/
	/*
		{
			start := time.Now()
			var counter Counter
			var wg sync.WaitGroup
			wg.Add(10)
			for i := 0; i < 10; i++ {
				go func() {
					defer wg.Done()
					for j := 0; j < 100000; j++ {
						counter.Lock()
						counter.Count++ // 非原子操作，好多计数会被吞掉
						counter.Unlock()
					}
				}()
			}
			wg.Wait()
			fmt.Println(time.Since(start))
			fmt.Println(counter.Count) // 最终的结果可能不是1000000，因为有些goroutine的计数被吞掉了，实测 300000 左右
		}
	*/
	/*
		{
			var mu sync.Mutex
			var wg sync.WaitGroup

			for i := 0; i < 5; i++ {
				wg.Add(1)
				go func(i int) {
					defer wg.Done()
					mu.Lock()
					defer mu.Unlock()
					fmt.Printf("Goroutine %d acquired the lock\n", i)
					time.Sleep(time.Second) // 模拟一些工作
				}(i)
			}

			wg.Wait()
		}

	*/
	{
		var mu sync.Mutex
		var wg sync.WaitGroup
		wg.Add(2)

		go func() {
			defer wg.Done()
			mu.Lock()
			fmt.Println("g1 get lock")
			time.Sleep(time.Second * 3)
		}()

		go func() {
			defer wg.Done()
			time.Sleep(time.Second * 2)
			//mu.Unlock()
			mu.Lock()
			fmt.Println("g2 get lock")
		}()
		wg.Wait()
	}

}
