package main

import (
	"fmt"
)

type Vertex struct {
	X int
	Y float64
}

// function can be passed as values
func computre(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func fn(x, y float64) float64 {
	return x + y
}

var m map[string]Vertex

func main() {
	fmt.Println(Vertex{1, 2})

	v := Vertex{4, 5}
	v.X = 67
	fmt.Println(v.X)

	// access an struct through a pointer
	p := &v
	// both of two lines below has same result
	fmt.Println((*p).X) // this line is very ugly
	fmt.Println(p.X)    // so we can use this one
	// and we can change struct parameters through a pointer

	// arrays
	var a [10]int // size of array is fixed
	a[0] = 2
	fmt.Println(a) // we can print the whole array
	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	// we can make slices of arrays
	var s []int = primes[1:4]
	fmt.Println(s)

	// * A slice does not store any data,
	//	 it just describes a section of an underlying array.
	// see the example below
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	slice1 := names[0:2]
	slice2 := names[1:3]
	fmt.Println(slice1, slice2)

	slice2[0] = "XXX"
	fmt.Println(slice1, slice2)
	fmt.Println(names)

	fmt.Printf("\n")

	var pow = []int{1, 2, 4, 8, 16, 32}
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}

	// the underline (_) is for skipping
	for _, v := range pow {
		fmt.Printf("%d", v)
	}

	// map
	// its some shit like dictionary
	// empty map be like
	// m = make(map[string]Vertex)

	m = map[string]Vertex{
		"google": Vertex{
			34, 34,
		},
	}
	m["Bell labs"] = Vertex{
		40, -74,
	}

	fmt.Println(m["Bell labs"])
	fmt.Println(m["google"])

	// delete a key from map
	delete(m, "google")
	elem, ok := m["google"]
	fmt.Println(elem, ok)

	fmt.Println(computre(fn))
}
