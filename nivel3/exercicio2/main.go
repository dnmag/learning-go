package main

import (
	"fmt"
)

func main() {
	for i := 65; i <= 90; i++ {
		fmt.Printf("%v\n", i)
		for x := 0; x < 3; x++ {
			fmt.Printf("\t%#U\n", i)
		}
	}
}
