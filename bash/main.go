package main

import (
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {

	tmpDir, _ := ioutil.TempDir("", "")
	defer os.RemoveAll(tmpDir)

	print(">>>>>", tmpDir)

	print("filepath : ")
	println(filepath.Join(">>>"))

	cmd := exec.Command("/bin/sh", "-c", "touch note.txt")
	_, err := cmd.Output()
	if err != nil {
		print(err.Error())
	}
}
