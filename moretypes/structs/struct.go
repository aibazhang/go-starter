package main

import "fmt"

type Vertex struct {
	X, Y int
}

var (
	v1 = Vertex{1, 2}
	v2 = Vertex{X: 1}
	v3 = Vertex{}
	p  = &Vertex{1, 2}
)

type language struct {
	Name     string
	LangType string
}

func NewLanguage(name string, langType string) *language {
	// l := new(Language)
	// l.Name = name
	// l.LangType = langType
	// return l

	// Simplified
	return &language{name, langType}
}

type DuckI interface {
	quack() string
	fly() string
}

type duck struct {
	Name string
	age  int
}

func NewDuck(name string, age int) DuckI {
	return &duck{name, age}
}

func (d *duck) quack() string {
	return "QUUAAAACCCCKKKKK!!!!!!"
}

func (d *duck) fly() string {
	return "I can NOT fly high"
}

func main() {
	// Coodinate
	// *p = Vertex{10000, 20000}
	// fmt.Println(v1, p, v2, v3)

	// Language
	// l := NewLanguage("Go", "Static")
	// fmt.Println("Name: " + l.Name)
	// fmt.Println("Language: " + l.LangType)

	// Duck
	d := NewDuck("Donald", 100)
	fmt.Println(d.quack())
	fmt.Println(d.fly())
}
