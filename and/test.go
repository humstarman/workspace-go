package main

import (
	"fmt"
)

func main() {
	xs := []bool{true, true, true}
	var x bool
	x = xs[0]
	for _, j := range xs[1:] {
		fmt.Println(j)
		x = x && j
	}
	fmt.Println(x)
}
