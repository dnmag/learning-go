package main

import (
	"fmt"
)

func main() {
	x := [][]string{
		[]string{"Bob", "Singer", "Caçar monstros"},
		[]string{"Dean", "Winchester", "Beber wisky e vadiar"},
		[]string{"Sam", "Winchester", "Estudar"},
	}
	for i, v := range x {
		fmt.Printf("%v\t%v\t%v\t%v\n", i, v[0], v[1], v[2])
	}
}
