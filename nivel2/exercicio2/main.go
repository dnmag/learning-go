package main

import (
	"fmt"
)

func main() {
	a := 1 == 1
	b := 5 != 16
	c := 10 <= 16
	d := 10 < 16
	e := 10 >= 16
	f := 10 > 16
	fmt.Println("A: ", a)
	fmt.Println("B: ", b)
	fmt.Println("C: ", c)
	fmt.Println("D: ", d)
	fmt.Println("E: ", e)
	fmt.Println("F: ", f)
}
