package shape

import (
	"fmt"
	"math"
	"reflect"
)

type shape interface {
	area() float64
	perimeter() float64
}

type polygon interface {
	getSides() int
}

type rectangle struct {
	width, height float64
	sides         int
}

type circle struct {
	radius float64
}

func (r rectangle) area() float64 {
	return r.height * r.width
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (r rectangle) perimeter() float64 {
	return 2*r.height + 2*r.width
}

func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func (r rectangle) getSides() int {
	return r.sides
}

func Calculate(s interface{}) {
	myType := reflect.TypeOf(s)
	fmt.Println(myType.Name())
	fmt.Println(myType.Kind())
	// fmt.Println(myType.Elem().Name())

	switch x := s.(type) {
	case circle:
		fmt.Printf("My type is %T\n", x)
	case rectangle:
		fmt.Printf("My type is %T\n", x)
	case int:
		fmt.Printf("My type is %T\n", x)
	default:
		fmt.Println("None of the above types")
	}
	fmt.Println("_________________________________")
}

func Run() {
	r := rectangle{width: 5, height: 8, sides: 4}
	c := circle{radius: 6}
	n := 10

	Calculate(r)
	Calculate(c)
	Calculate(n)

	Calculate(&r)
	Calculate(&c)
	Calculate(&n)
}
