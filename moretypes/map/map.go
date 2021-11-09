package main

import (
	"fmt"
	"strings"
)

type Vertex struct {
	Lat, Long float64
}

// var m = map[string]Vertex{
// 	"Bell Labs": {40.68433, -74.39967},
// 	"Google":    {37, -122},
// }

func MutateMap() {
	m := make(map[string]int)
	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)

	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	v, ok = m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)

}

func WordCount(s string) map[string]int {
	wordArray := strings.Split(s, " ")
	wordCount := make(map[string]int)
	for _, v := range wordArray {
		wordCount[v] += 1
	}
	return wordCount
}

func main() {
	// fmt.Println(m)
	// MutateMap()
	fmt.Println(WordCount("ha a edede ha a test"))
}
