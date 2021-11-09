package main

import "fmt"

func arrays() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"

	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
}

func slices() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[1:4]
	fmt.Println(s)
}

func slices_share_underlying_arrary() {
	names := []string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)
}

func slice_iteral() {
	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func slice_length_and_capacity() {

	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// Slice the slice to give it zero length.
	s = s[:0]
	printSlice(s)

	// Extend its length.
	s = s[:4]
	printSlice(s)

	s = s[:5]
	printSlice(s)

	// Drop its first two values.
	s = s[2:]
	printSlice(s)
}

func for_range() {
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
	pow = append(pow, 256, 512)
	for i, v := range pow {
		fmt.Printf("2**%v = %v\n", i, v)
	}
}

func main() {
	// arrays()
	// slices()
	// slices_share_underlying_arrary()
	// slice_iteral()
	// slice_length_and_capacity()
	for_range()
}
