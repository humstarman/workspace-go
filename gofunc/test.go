package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func trace(name int) {
	for i := 0; i < 10; i++ {
		fmt.Printf("name: %v, ret: %v\n", name, i)
		time.Sleep(time.Second)
	}
}

func main() {
	n := 4
	runtime.GOMAXPROCS(n)
	var wg sync.WaitGroup
	wg.Add(n)

	for i := 0; i < n; i++ {
		go func(i int) {
			defer wg.Done()
			trace(i)
		}(i)
	}
	wg.Wait()
}
