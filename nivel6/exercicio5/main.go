package main

import (
	"fmt"
	"math"
)

type square struct {
	width  float64
	height float64
}

func (s square) area() float64 {
	return s.width * s.height
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return 2 * math.Pi * c.radius
}

type figure interface {
	area() float64
}

func info(f figure) float64 {
	return f.area()
}

func main() {
	x := square{10, 20}
	y := circle{20}
	fmt.Println(info(x))
	fmt.Println(info(y))
}
