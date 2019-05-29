package main

import (
	"fmt"
	"crypto/sha1"
	"io"
)

func makeHash(data string) string {
	h := sha1.New()
	io.WriteString(h,data)
	return fmt.Sprintf("%x",h.Sum(nil))
}

func main() {
	str0 := "hi"
	str1 := "hello"
	h0 := makeHash(str0)
	h1 := makeHash(str1)
	fmt.Println(h0)
	fmt.Println(h1)
}
