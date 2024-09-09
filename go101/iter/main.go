// package main
package main

import (
	"fmt"
	"iter"
)

func main() {
	for i := range stat(10) {
		fmt.Println(i)
	}

}

func stat(v int) iter.Seq[int] {
	return func(yield func(int) bool) {
		for i := v; i >= -1; i-- {
			if !yield(i) {
				return
			}
		}
	}
}
