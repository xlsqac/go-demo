// package main
package main

import (
	"fmt"
	"time"
)

type T struct {
	n int
}

func (t *T) pn() {
	fmt.Println(t.n)
}

func (t T) pn1() {
	fmt.Println(t.n)
}

func main() {
	t := T{1}
	t1 := &T{2}
	t.pn()
	t.pn1()
	t1.pn()
	t1.pn1()

	// 个人理解拷贝出来的 timer 没有使用 NewTimer or AfterFunc
	// 而文档中的说明是，一个 timer 必须使用 NewTimer or AfterFunc 来初始化
	// 这里的 *time.NewTimer() 实际上进行了一个对 timer 的拷贝，生成了一个新的 timer
	// 而这个新的 timer 实际是没有初始化的，所以 Stop 的时候出问题了
	// 在 linux 环境下是会一直卡主
	illegalTimerCopy := *time.NewTimer(time.Second)
	illegalTimerCopy.Stop() // block for ever
}
