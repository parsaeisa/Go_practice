package main

import "regexp"

func main() {
	// learn regex notations
	regex, err := regexp.Compile("foo.*")
	if err != nil {
		panic(err)
	}

	matched := regex.MatchString("foo.heasdfa")
	print(matched)
}
