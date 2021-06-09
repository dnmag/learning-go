package main

import (
	"fmt"
)

func main() {
	x := 55
	fmt.Printf("A variável x possui o valor %v e está no endereço %v\n", x, &x)
}
