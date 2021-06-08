package main

import (
	"fmt"
)

func main() {
	x := map[string][]string{
		"Singer_Bob":      []string{"Caçar monstros", "Beber wisky e vadiar"},
		"Winchester_Dean": []string{"Caçar monstros", "Beber wisky e vadiar"},
		"Winchester_Sam":  []string{"Caçar monstros", "Estudar"},
	}

	x["Winchester_Mary"] = []string{"Caçar monstros", "Matar demonios"}

	delete(x, "Singer_Bob")

	for key, v := range x {
		fmt.Println(key)
		for i, hobbie := range v {
			fmt.Printf("\t%v \t%v\n", i, hobbie)
		}
	}
}
