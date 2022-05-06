package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func subtract(x, y int) int {
	return x - y
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum + 2
	y = sum - 2
	return
}

func RunFunctions() {
	x := 2
	y := 3
	fmt.Println("Sum ", x, y, " => ", add(x, y))
	fmt.Println("Subtract ", x, y, " => ", subtract(x, y))

	a, b := swap("Hello", "World")
	fmt.Println("Swap =>", a, b)
	fmt.Println(swap("Hi", "There"))

	// fmt.Println("Split =>", split(2))
	fmt.Println(split(2))
}
