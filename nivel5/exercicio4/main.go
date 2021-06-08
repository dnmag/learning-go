package main

import (
	"fmt"
)

func main() {
	x := struct {
		meuMap   map[string]int
		meuSlice []string
	}{
		meuMap: map[string]int{
			"teste":  1,
			"teste2": 2,
		},
		meuSlice: []string{
			"teste1", "teste2",
		},
	}

	fmt.Println(x)
}
