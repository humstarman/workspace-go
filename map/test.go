package main

import (
	"fmt"
)

func main() {
	m := make(map[string]string)
	m["a"] = "A"
	m["b"] = "B"
	for k, v := range m {
		fmt.Printf("%v: %v\n", k, v)
	}
}
