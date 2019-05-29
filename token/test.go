package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"log"
	"strings"
)

func main() {
	fn, err := os.OpenFile("/tmp/token.csv", os.O_RDONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer fn.Close()
	b, err := ioutil.ReadAll(fn)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(b))
	strs := strings.Split(strings.Trim(string(b),"\n"),",")
	fmt.Println(strs)
	bootstrapToken := strs[0]
	fmt.Println(bootstrapToken)
}
