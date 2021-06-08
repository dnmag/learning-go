package main

import (
	"fmt"
)

const (
	bornYear    = 1998
	currentYear = 2021
)

func main() {

	countYear := bornYear
	for {
		if countYear <= currentYear {
			fmt.Printf("%v\n", countYear)
			countYear++
		} else {
			break
		}
	}
}
