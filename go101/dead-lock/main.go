// package main
package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.NewTimer(time.Minute * 5)
	defer t.Stop()
	select {
	case <-t.C:
		fmt.Println("time out", time.Now().Unix())
	}

}
