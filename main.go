package main

import (
	"fmt"
	"os"
	"runtime"
	"runtime/pprof"
	"sync"
	// _ "go.uber.org/automaxprocs"
)

func cpuIntensiveTask(wg *sync.WaitGroup, id int) {
	defer wg.Done()
	var x int
	for i := 0; i < 1e9; i++ {
		x += i
	}
	fmt.Printf("Goroutine %d done\n", id)
}

func main() {
	n := runtime.GOMAXPROCS(2)
	fmt.Println("Proc: ", n)
	runtime.GOMAXPROCS(n)

	f, err := os.Create("cpu.prof")
	if err != nil {
		fmt.Println("could not create CPU profile:", err)
		return
	}
	defer f.Close()
	if err := pprof.StartCPUProfile(f); err != nil {
		fmt.Println("could not start CPU profile:", err)
		return
	}
	defer pprof.StopCPUProfile()

	var wg sync.WaitGroup

	for i := 1; i <= 2; i++ {
		wg.Add(1)
		go cpuIntensiveTask(&wg, i)
	}

	wg.Wait()
}
