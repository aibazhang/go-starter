package main

import (
	"fmt"
)

type I interface {
	M()
}

type T struct {
	S string
}

// Type T implements the interface I
func (t *T) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println((t.S))
}

type F float64

func (f F) M() {
	fmt.Println(f)
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}

// Type switches
func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

func main() {
	/*
		var i I

		i = &T{"Hello"}
		describe(i)
		i.M()

		i = F(math.Pi)
		describe(i)
		i.M()

		// Interface values with nil underlying values
		var t *T
		i = t
		describe(i)
		i.M()
	*/

	// Nil interface values
	// var i I
	// describe(i)
	// i.M()
	do(21)
	do("hello")
	do(true)
}
