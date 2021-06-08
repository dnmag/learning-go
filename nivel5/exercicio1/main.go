package main

import (
	"fmt"
)

type Pessoa struct {
	nome                    string
	sobrenome               string
	saboresFavoritosSorvete []string
}

func main() {
	pessoa1 := Pessoa{
		nome:      "Joaquim",
		sobrenome: "Alcantara",
		saboresFavoritosSorvete: []string{
			"Morango",
			"Baunilha",
		},
	}
	pessoa2 := Pessoa{
		nome:      "Luisa",
		sobrenome: "Ferreira",
		saboresFavoritosSorvete: []string{
			"Chocolate",
			"Amexa",
		},
	}

	fmt.Println(pessoa1.nome, "gosta dos seguintes sabores:")
	for _, v := range pessoa1.saboresFavoritosSorvete {
		fmt.Printf("\t%v\n", v)
	}

	fmt.Println(pessoa2.nome, "gosta dos seguintes sabores:")
	for _, v := range pessoa2.saboresFavoritosSorvete {
		fmt.Printf("\t%v\n", v)
	}
}
