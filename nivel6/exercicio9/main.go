package main

import (
	"fmt"
)

func usaCallback(callback func()) {
	callback()
}

func main() {
	x := func() {
		fmt.Println("Essa é uma função que está sendo passada como callback!!!")
	}
	usaCallback(x)
}
