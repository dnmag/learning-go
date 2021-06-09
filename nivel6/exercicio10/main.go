package main

import (
	"fmt"
)

func contadorFunc() func() {
	contador := 0
	return func() {
		fmt.Println("O contador está em:", contador)
		contador++
	}
}

func main() {
	x := contadorFunc()
	x()
	x()
	x()
	y := contadorFunc()
	y()
	y()
	y()
}
