package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	time.Sleep(1 * time.Second)
	t := time.Now()
	fmt.Println(t.Sub(start))
}
