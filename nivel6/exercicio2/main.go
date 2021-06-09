package main

import (
	"fmt"
)

func main() {
	x := []int{54, 22, 34, 66, 89, 23, 11, 54, 65, 34}

	fmt.Println(variadicSum(x...))
	fmt.Println(sliceSum(x))
}

func variadicSum(x ...int) int {
	return sliceSum(x)
}

func sliceSum(x []int) int {
	total := 0
	for _, v := range x {
		total += v
	}
	return total
}
