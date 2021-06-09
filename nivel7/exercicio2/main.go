package main

import (
	"fmt"
)

type Pessoa struct {
	nome      string
	sobrenome string
}

func (p *Pessoa) mudeMe() {
	(*p).nome = "Nome Alteado"
}

func main() {
	pessoa1 := Pessoa{
		nome:      "Joaquim",
		sobrenome: "Alcantara",
	}

	fmt.Println(pessoa1)
	pessoa1.mudeMe()
	fmt.Println(pessoa1)
}
