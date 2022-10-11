package main

import (
	"fmt"
	"regexp"
)

func main() {
	// learn regex notations
	regex, err := regexp.Compile("foo.*")
	if err != nil {
		panic(err)
	}

	matched := regex.MatchString("foogg")
	println(matched)

	re := regexp.MustCompile("foo.?")
	fmt.Printf("%q\n", re.Find([]byte("seafood foollee")))
}
