package main

import (
	"fmt"
)

type pessoa struct {
	nome      string
	sobrenome string
	idade     int
}

func (p pessoa) nomeCompletoEIdade() {
	fmt.Printf("Meu nome é %v %v e tenho %v anos\n", p.nome, p.sobrenome, p.idade)
}

func main() {
	x := pessoa{"Jubileu", "da Silva", 43}
	x.nomeCompletoEIdade()
}
