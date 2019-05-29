package main

import (
	"fmt"
)

const (
	v1 = `1
2
`
	v2 = `x
y
z
`
)

func main() {
	fmt.Println(v1)
	fmt.Println(v2)
	var s string
	s += string(v1)
	s += "a\n"
	s += "b\n"
	s += "c\n"
	s += string(v2)
	fmt.Println(s)
}
