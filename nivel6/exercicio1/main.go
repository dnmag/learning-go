package main

import (
	"fmt"
)

func main() {
	fmt.Println(returnsInt())
	fmt.Println(returnsIntAndString())
}

func returnsInt() int {
	return 546
}

func returnsIntAndString() (int, string) {
	return 679, "Mil linhas de código"
}
