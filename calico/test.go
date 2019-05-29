package main

import (
	"fmt"
)

const test = `hi
hello
test
`

func main() {
	str := string(test)
	str += "`cat /tmp/log | tr '\\n' ''`"
	fmt.Println(str)
	str0 := "/tmp/data"
	fmt.Println(str0)
	x := "manifest"
	str0 += "/"
	str0 += x
	str0 += "\n"
	fmt.Println(str0)

	manifest := "manifest"
	content := "" 
	content += str 
	content += "\n" 
	content += str0 
	content += "\n" 
	content += fmt.Sprintf("MANIFEST=./%v\n",manifest)
	content += str 
	content += "\n" 
	content += str0 
	content += "\n" 
	fmt.Println(content)
}
