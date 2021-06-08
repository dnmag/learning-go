package main

import (
	"fmt"
)

type veiculo struct {
	portas int
	cor    string
}

type caminhonete struct {
	veiculo           veiculo
	tracaoQuatroRodas bool
}

type sedan struct {
	veiculo    veiculo
	modeloLuxo bool
}

func main() {
	c := caminhonete{
		veiculo: veiculo{
			portas: 4,
			cor:    "Preta",
		},
		tracaoQuatroRodas: true,
	}

	s := sedan{
		veiculo: veiculo{
			portas: 4,
			cor:    "Prata",
		},
		modeloLuxo: true,
	}

	fmt.Println(c)
	fmt.Println(s)
	fmt.Println(c.veiculo.cor)
	fmt.Println(s.modeloLuxo)
}
