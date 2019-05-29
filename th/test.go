package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func print(i int) {
	if i == 0 {
		fmt.Printf("%v here\n", i)
		return
	}
	for j := 0; j < 10; j++ {
		fmt.Printf("%v: %v\n", i, j)
		time.Sleep(time.Second)
	}
	return
}

func main() {
	n := 5
	runtime.GOMAXPROCS(n)
	var wg sync.WaitGroup
	wg.Add(n)
	go func(i int) {
		defer wg.Done()
		print(i)
	}(0)
	for i := 1; i < n; i++ {
		go func(i int) {
			defer wg.Done()
			print(i)
		}(i)
	}
	wg.Wait()
	fmt.Println("end")
}
