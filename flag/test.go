package main

import (
	"flag"
	"log"
)

var (
	docker = flag.Bool("d", false, "Force to boot as a new node")
)

func init() {
	flag.Parse()
}

func main() {
	log.Println(*docker)
}
