package main

import (
	"fmt"
)

func retornaUmaFuncao() func() {
	return func() {
		fmt.Println("Essa é uma função que foi retornada de outra função!!!")
	}
}

func main() {
	x := retornaUmaFuncao()
	x()
}
