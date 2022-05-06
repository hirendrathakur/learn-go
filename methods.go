package main

import (
	"fmt"
	"math"
)

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
func RunMethods() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
}
