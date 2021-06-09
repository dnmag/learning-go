package main

import (
	"fmt"
)

func main() {
	defer fmt.Println("Chamei esse primeiro")
	fmt.Println("Mas esse executou primeiro")
}
