package main

import (
	//"io/ioutil"
	"bufio"
	"io"
	"log"
	"os"
	"regexp"
)

func containCheck(b []byte) bool {
	pattern := []string{"allow.+", "server.+iburst", "haha"}
	for _, p := range pattern {
		res, _ := regexp.Match(p, b)
		if res {
			return res
		}
	}
	return false
}

func main() {
	path := "/tmp/test.conf"
	var last []byte
	fn, err := os.OpenFile(path, os.O_RDONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer fn.Close()

	br := bufio.NewReader(fn)
	for {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		if "" != string(a) {
			last = a
		}
	}
	// 1 check
	log.Println(containCheck(last))
}
