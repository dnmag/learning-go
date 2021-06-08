package main

import (
	"fmt"
)

func main() {
	horaAtual := 10

	switch {
	case horaAtual > 8 && horaAtual < 12:
		fmt.Println("Estou trabalhando")

	case horaAtual > 12 && horaAtual < 13:
		fmt.Println("Estou no horário de almoço \\o/")

	case horaAtual > 13 && horaAtual < 17:
		fmt.Println("Trbalhando, segundo turno")

	default:
		fmt.Println("Provavelmente estou estudando ou dormindo!!!zzz")
	}
}
