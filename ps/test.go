package main

import (
	"fmt"
	"io/ioutil"
	//"log"
	//"math/rand"
	"os/exec"
	//"strconv"
	//"strings"
)

func execCmd(strCommand string) (string, error) {
	cmd := exec.Command("/bin/bash", "-c", strCommand)

	stdout, _ := cmd.StdoutPipe()
	if err := cmd.Start(); err != nil {
		fmt.Println("Execute failed when Start:" + err.Error())
		return "", err
	}

	out_bytes, _ := ioutil.ReadAll(stdout)
	stdout.Close()

	if err := cmd.Wait(); err != nil {
		fmt.Println("Execute failed when Wait:" + err.Error())
		return "", err
	}
	str := string(out_bytes)
	return str, nil
}

func main() {
	componet := "/usr/sbin/VBoxService"
	cmd := fmt.Sprintf("ps aux | grep -v grep | grep \"%v\"",componet)
	fmt.Println(cmd)
	res, _ := execCmd(cmd)
	fmt.Println(res)
	fmt.Println(res == "")
}
