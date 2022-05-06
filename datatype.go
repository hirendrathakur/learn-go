package main

import "fmt"

type Vertex struct {
	X, Y float64
}

func RunDataTypes() {
	//Vertex definition
	fmt.Println(Vertex{1, 2})

	//array definition
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	var s []int = primes[1:4]
	fmt.Println(s)
}
