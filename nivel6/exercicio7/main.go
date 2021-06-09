package main

import (
	"fmt"
)

func main() {
	x := func() {
		fmt.Println("Essa é uma função atribuida a uma variável!!!")
	}
	x()
}
