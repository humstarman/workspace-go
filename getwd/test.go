package main

import (
	"fmt"
	"os"
	//"log"
	"path/filepath"
)

func main() {
	path, _ := filepath.Abs(filepath.Dir(os.Args[0])) 
	fmt.Println(path)
	fmt.Println(os.Args[0])
}
