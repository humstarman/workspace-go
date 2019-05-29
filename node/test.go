package main

import (
	"fmt"
)

func main() {
	a := []string{"1", "2"}
	var b []string
	for _, i := range a {
		fmt.Println(i)
	}
	fmt.Println(len(b))
}
