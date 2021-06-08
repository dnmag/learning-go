package main

import (
	"fmt"
)

func main() {
	x := [10]int{54, 22, 34, 66, 89, 23, 11, 54, 65, 99}

	fmt.Println(x[0:3])
	fmt.Println(x[4:])
	fmt.Println(x[1:7])
	fmt.Println(x[2:9])
	fmt.Println(x[2 : len(x)-1])
}
