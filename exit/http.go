package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
)

var (
	port = flag.String("p", "8080", "Specify the port")
)

func init() {
	flag.Parse()
}

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

func handleExitRequest(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		ret := make(map[string]string)
		ret["PodIP"] = os.Getenv("POD_IP")
		b, err := json.Marshal(ret)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		io.WriteString(w, string(b))
		os.Exit(1)
	}
}

func main() {
	http.HandleFunc("/", handleExitRequest)
	addr := fmt.Sprintf("0.0.0.0:%v", *port)
	http.ListenAndServe(addr, nil)
}
