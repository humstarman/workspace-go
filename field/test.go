package main

import (
	"fmt"
	"reflect"
	"log"
	"os/exec"
	"io/ioutil"
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

type kubeclusterstatus struct {
	Summary string           `json:"summary"`
	Status  []kubenodestatus `json:"status"`
}

type kubenodestatus struct {
	Role                  string 
	KubeApiserver         string
	KubeControllerManager string
	KubeScheduler         string
	Kubelet               string
	KubeProxy             string
	Docker                string
}

func reportKubeStatus() kubeclusterstatus {
	ret := new(kubeclusterstatus)
	return *ret
}

func getKubeNodeStatus() kubenodestatus {
        ret := new(kubenodestatus)
        r := reflect.ValueOf(ret).Elem()
        componets := make(map[string]string)
        componets["KubeApiserver"] = "kube-apiserver"
        componets["KubeControllerManager"] = "kube-controller-manager"
        componets["KubeScheduler"] = "kube-scheduler"
        componets["Kubelet"] = "kubelet"
        componets["KubeProxy"] = "kube-proxy"
        componets["Docker"] = "docker"
        for f, u := range componets {
                cmd := fmt.Sprintf("systemctl is-active %v", u)
		fmt.Println(cmd)
                v, err := execCmd(cmd)
                if err != nil {
                        log.Println(err)
                }
                r.FieldByName(f).Set(reflect.ValueOf(v))
        }
	ret.Role = "test"
        return *ret
}

func main() {
	k := getKubeNodeStatus() 
	fmt.Printf("%v\n",k)
	var x [3]bool
	fmt.Println(x)
}
