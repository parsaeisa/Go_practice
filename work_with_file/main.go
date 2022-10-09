package main

import (
	"bytes"
	"io/ioutil"
)

const (
	OldName = "parsa1"
	NewName = "parsa2"
)

func main() {
	err := changeNameInFile("./test.txt", "./dest.txt")
	if err != nil {
		println(err.Error())
	}
}

func changeNameInFile(templateFilePath, destinationFilePath string) error {
	content, err := ioutil.ReadFile(templateFilePath)
	if err != nil {
		return err
	}

	// read about what is files
	newContent := bytes.Replace(content, []byte(OldName), []byte(NewName), -1)

	err = ioutil.WriteFile(destinationFilePath, newContent, 0644)
	if err != nil {
		return err
	}

	return nil
}

// problem :
// try to copy some files ( which might be in some directories ) , using ox/exec package
// change some name in them
// something like jump
