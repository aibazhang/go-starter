package main

import (
	"fmt"
)

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return x
	}
}

func PosNeg() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(pos(i), neg(-2*i))
	}
}

func fibonacci(x int) func() int {
	first, second := 0, 1
	return func() int {
		ret := first
		first, second = second, first+second
		return ret
	}
}

func main() {
	// PosNeg()

}
