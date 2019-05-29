package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
)

func containCheck(b []byte) {
	pattern := []string{"allow.+", "server.+iburst", "haha"}
	for _, p := range pattern {
		log.Println(p)
		res, _ := regexp.Compile(p)
		log.Println(res)
		fmt.Println(res.FindAllString(string(b), -1))

	}
}

func main() {
	path := "/tmp/test.conf"
	fn, err := os.OpenFile(path, os.O_RDONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer fn.Close()
	b, err := ioutil.ReadAll(fn)
	if err != nil {
		log.Println(err)
	}
	// 1 check
	containCheck(b)
}
