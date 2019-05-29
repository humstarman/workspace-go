package main

import (
	"fmt"
	"log"
	"time"
)

var (
	ch = make(chan int)
)

func trace() {
	for i := 0; i < 10; i++ {
		log.Println(i)
		ch <- i
		time.Sleep(time.Second)
	}
}

func recv() {
	for {
		i := <-ch
		fmt.Println(i)
	}
}

func main() {
	go recv()
	trace()
}
