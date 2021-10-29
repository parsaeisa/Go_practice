package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

// Constants cannot be declared using the := syntax.
const parsa = 25

func add(x int, y int) int { // it can be func add (x, y int ) int
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) { // returned values can be named
	x = sum * 4 / 9
	y = sum - x
	return
}

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

var c, python, java bool

func main() {
	fmt.Println(math.Pi)
	fmt.Println(add(42, 14))

	a, b := swap("hello", "world")
	fmt.Println(a, b)

	var i int
	var j, k int = 1, 2                   // initializing with var
	fmt.Println(j, k, i, c, python, java) // var is for defining list of variables

	//Outside a function, every statement begins with a keyword (var, func, and so on) and so the := construct is not available.
	h := 4
	fmt.Println(h)

	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
}
