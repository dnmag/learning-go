package main

import (
	"fmt"
)

func main() {
	x := [10]int{54, 22, 34, 66, 89, 23, 11, 54, 65, 34}

	for _, v := range x {
		fmt.Println(v)
	}
	fmt.Printf("%T\n", x)
}
