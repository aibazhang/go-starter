package main

import "fmt"

// Deferred calls are executed in last-in-first-out
func main() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("done")
}
