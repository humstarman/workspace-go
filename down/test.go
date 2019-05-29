package main

import (
	"fmt"
)

const (
	cfsslVersion="R1.2"
)


type downlink struct {
	url  string
	path string
}

func (d *downlink) check() bool {
	return checkFileExist(d.path)
}

func (d *downlink) download() (string, error) {
	if d.check() {
		return "Already existed", nil
	}
	cmd := new(linuxcmd)
	var arg string
	if d.path == "" {
		arg = ""
	} else {
		arg = fmt.Sprintf("-O %v", d.path)
	}
	cmd.set(fmt.Sprintf("wget %v", d.url), arg)
	return cmd.run()
}

func (d *downlink) set(durl, path string) error {
	d.url = durl
	d.path = path
	return nil
}


func main(){
        urls := []string{fmt.Sprintf("https://pkg.cfssl.org/%v/cfssl_linux-amd64", cfsslVersion), fmt.Sprintf("https://pkg.cfssl.org/%v/cfssljson_linux-amd64", cfsslVersion), fmt.Sprintf("https:/
/pkg.cfssl.org/%v/cfssl-certinfo_linux-amd64", cfsslVersion)}
        paths := []string{"/tmp/cfssl", "/tmp/cfssljson", "/tmp/cfssl-certinfo"}
        for i := 0; i < 3; i++ {
                d := new(downlink)
                d.set(urls[i], paths[i])
                d.download()
        }
        return nil

}
