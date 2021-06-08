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
	pessoas := map[string]Pessoa{}
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
	pessoas["Alcantara"] = pessoa1
	pessoas["Ferreira"] = pessoa2

	for _, p := range pessoas {

		fmt.Println(p.nome, "gosta dos seguintes sabores:")
		for _, v := range p.saboresFavoritosSorvete {
			fmt.Printf("\t%v\n", v)
		}
	}
}
