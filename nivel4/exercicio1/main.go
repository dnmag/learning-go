package main

import (
	"fmt"
)

func main() {
	x := [5]int{54, 22, 34, 66, 89}

	for _, v := range x {
		fmt.Println(v)
	}
	fmt.Printf("%T\n", x)
}
