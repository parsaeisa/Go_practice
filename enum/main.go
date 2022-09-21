package main

// this is a pattern which makes and enum of fixed numbers and strings
type Fruit int

const (
	Apple Fruit = iota
	Orange
	Banana
)

func (f Fruit) String() string {
	return [...]string{
		"Apple",
		"Orange",
		"Banana",
	}[f]

	// Question : what is ... in golang
}

func main() {
	println(Orange.String())
}
