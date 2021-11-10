package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

type MyFloat float64

// Declare a method on ono-struct types, too
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

// Pointer receivers (methods often need to modify their receiver
// pointer recrivers are more common than value receivers)
func (v *Vertex) Scale(f float64) {
	v.X *= f
	v.Y *= f
}

func ScaleFunc(v *Vertex, f float64) {
	v.X *= f
	v.Y *= f
}

func main() {
	// v := Vertex{3, 4}
	// fmt.Println(v.Abs())

	// f := MyFloat(-math.Sqrt2)
	// fmt.Println(f.Abs())

	// v := Vertex{3, 4}
	// v.Scale(2)
	// ScaleFunc(&v, 10)
	// fmt.Println(v)

	// p := &Vertex{4, 3}
	// p.Scale(2)
	// ScaleFunc(p, 10)
	// fmt.Println(p)

	v := &Vertex{3, 4}
	fmt.Printf("Before scaling: %+v, Abs: %v\n", v, v.Abs())
	v.Scale(5)
	fmt.Printf("After scaling: %+v, Abs: %v\n", v, v.Abs())
}
