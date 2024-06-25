// package main
// https://www.yinkai.cc/post/beea4a4c31ba01d748a7e5186cdcfb11
// 利用 context 控制 goroutine
// 按下 ctrl-c 后，利用 context 把 g 全部关闭
package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func doRequest(ctx context.Context, id int) error {
	req, _ := http.NewRequestWithContext(ctx, "GET", "http://baidu.com", nil)
	fmt.Printf("id: %d, method: %s, url: %s\n", id, req.Method, req.URL)

	select {
	case <-time.After(2 * time.Second):
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}

}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// 监听 ctrl+c 信号
	signChan := make(chan os.Signal, 1)
	signal.Notify(signChan, syscall.SIGINT)

	var wg sync.WaitGroup

	// 5 个协程来发送请求
	n := 5
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			err := doRequest(ctx, id)
			if err != nil {
				fmt.Printf("请求 #%d 被取消或出现错误: %v\n", id, err)
			} else {
				fmt.Printf("请求 #%d 完成", id)
			}
		}(i)
	}

	go func() {
		<-signChan
		fmt.Println("接收到Ctrl+C，正在取消所有请求...")
		cancel()
	}()

	wg.Wait()
	fmt.Println("所有请求处理完毕，程序退出")
}
