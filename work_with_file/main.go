package main

import (
	"bytes"
	"io/ioutil"
	"path"
)

const (
	OldName = "parsa1"
	NewName = "parsa2"
)

func main() {
	//err := changeNameInFile("./test.txt", "./dest.txt")
	//if err != nil {
	//	println(err.Error())
	//}

	err := printFileNamesInDir("./../", "")
	if err != nil {
		panic(err)
	}
}

// how we define optional parameter with default value
func printFileNamesInDir(s, whitespace string) error {
	entries, err := ioutil.ReadDir(s)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		if entry.IsDir() {
			print(">>>>>>")
			err = printFileNamesInDir(path.Join(s, entry.Name()), whitespace+"  ")
			if err != nil {
				return err
			}
		} else {
			print(whitespace)
			println(entry.Name())
		}

	}

	return nil
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
