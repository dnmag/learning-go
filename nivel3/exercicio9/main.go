package main

import (
	"fmt"
)

func main() {
	esporteFavorito := "Mergulho"

	switch esporteFavorito {
	case "Mergulho":
		fmt.Println("Meu esporte favorito é Mergulho")
	case "Volei":
		fmt.Println("Meu esporte favorito é Volei")
	case "Natação":
		fmt.Println("Meu esporte favorito é Natação")
	case "Futebol":
		fmt.Println("Meu esporte favorito é Futebol")
	case "ESports":
		fmt.Println("Meu esporte favorito é ESports")
	default:
		fmt.Println("Sou muito sedentário!")
	}
}
