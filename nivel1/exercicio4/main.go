package main

import "fmt"

type NewType int

var x NewType

func main() {
	fmt.Printf("%v\n", x)
	fmt.Printf("%T\n", x)

	x = 42
	fmt.Println(x)
}
