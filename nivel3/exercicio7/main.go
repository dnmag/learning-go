package main

import (
	"fmt"
)

func main() {
	x := 33
	if x >= 133 {
		fmt.Println("X é maior ou igual a 133")
	} else if x < 133 && x > 50 {
		fmt.Println("X é menor que 133 e maior que 50")
	} else {
		fmt.Println("X é menor que 50")
	}
}
