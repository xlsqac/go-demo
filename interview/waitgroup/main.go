// package main
package main

import (
	"fmt"
	"sync"
)

// https://mp.weixin.qq.com/s/W-BOtDkKO-JaeQKqD_wU1A
// 问题 4，输出什么，什么都不会输出，wg 没有 Wait 进程直接退出了。
// 如果在代码最后加了 wg.Wait() 会随机输出 0 到 9 之间的数组。
// 在 golang version 1.22 之前？可能会输出 10 个 9。
// 因为使用的是外层的 i，而 i 的值会循环变化，但是引用不会变化，最后等 gouroutine 开始执行时，i 的值已经是 9 了。
func main() {
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			fmt.Println(i)
			wg.Done()
		}(&wg)
	}
	// wg.Wait()
}
