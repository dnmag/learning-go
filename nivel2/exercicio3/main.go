package main

import (
	"fmt"
)

const (
	a float64 = 10
	b         = 30
)

func main() {
	fmt.Printf("Typed: %v\t%T\nUntyped: %v\t%T\n", a, a, b, b)
}
