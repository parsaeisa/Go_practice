package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	x, y float64
}

func (v Vertex) Abs() float64 {

	// after doing this v has an Abs function
	// as one of its properties
	// instead of vertex we can put any other types
	return math.Sqrt(v.x*v.x + v.y*v.y)
}

// There are two reasons to use a pointer receiver :

// * The first is so that the method can modify
//	 	the value that its receiver points to.
// * The second is to avoid copying the value on each method call.
// 		This can be more efficient if the receiver is a large struct, for example.
func (v *Vertex) Scale(f float64) {
	v.x = v.x * f
	v.y = v.y * f
}

type Abser interface {
	Abs() float64
}

type MyFloat float64

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())

	v.Scale(5) // instead of (*v).Scale(5)
	fmt.Printf("After scaling: %+v, Abs: %v\n", v, v.Abs())

	// any type can implement an interface
	var a Abser
	f := MyFloat(12.333)
	a = f

	fmt.Println(a.Abs())
}

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}
