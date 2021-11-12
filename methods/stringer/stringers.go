package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v year)", p.Name, p.Age)
}

type IPAddr [4]byte

// TODO: Add a "String() string" method to IPAddr.
func (ip IPAddr) String() string {
	r := ""
	for _, v := range ip {
		if len(r) > 0 {
			r += "."
		}
		r += strconv.Itoa(int(v))
	}
	return r
}

func main() {
	// a := Person{"Arthur Dent", 42}
	// z := Person{"Zaphod Beeblebrox", 9001}
	// fmt.Println(a, z)

	// Int Array to String Array
	a := [6]int{2, 3, 5, 7, 11, 13}
	r1 := strings.Trim(strings.Join(strings.Split(fmt.Sprint(a), " "), "."), "[]")
	fmt.Println(a)
	fmt.Println(r1)

	ip := IPAddr{8, 8, 8, 8}
	fmt.Println(ip)
}
