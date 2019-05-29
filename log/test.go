package main

import (
	//"fmt"
	"log"
	"os"
	"errors"
)

func main() {
	path := "/tmp/test.log"
	f, err := os.Create(path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	logger := log.New(f, "", log.LstdFlags)
	logger.Println("test")
	logger.Println("test")
	logger.Println("test")
	err = errors.New("test error")
	logger.Fatal(err)
}
