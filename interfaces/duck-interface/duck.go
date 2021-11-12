package main

import (
	"bytes"
	"fmt"
)

// Duck
type DuckI interface {
	quack() string
}

type duck struct{}

func NewDuck() DuckI {
	return &duck{}
}

func (d *duck) quack() string {
	return "QUUAAAACCCCKKKKK!!!!!!"
}

// Self Calling Duck
type selfCallingDuck struct {
	name string
}

func NewNameDuck(name string) DuckI {
	return &selfCallingDuck{name}
}

func (d *selfCallingDuck) quack() string {
	return d.name
}

// Farmer
type FarmerI interface {
	Breed() string
}

type farmer struct {
	ducks []DuckI
}

func NewFarmer(ducks ...DuckI) FarmerI {
	return &farmer{ducks}
}

func (f *farmer) Breed() string {
	var b bytes.Buffer
	for i, d := range f.ducks {
		fmt.Fprintf(&b, "#%d %s\n", i, d.quack())
	}
	return b.String()
}

func main() {
	duck1 := NewDuck()
	duck2 := NewNameDuck("Donald")
	farmer := NewFarmer(duck1, duck2)
	fmt.Println(farmer.Breed())
}
