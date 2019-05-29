package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	n := fmt.Sprintf("%06v", rand.Int())
	fmt.Println(n)
	n = fmt.Sprintf("%02v", rand.Int())
	fmt.Println(n)
	n = fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))
	fmt.Println(n)
	n = fmt.Sprintf("%08v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))
	fmt.Println(n)
}
