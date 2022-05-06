package main

import "fmt"
import "math"

func RunLoops() {
	forLoop()
	whileLoop()
	ifStatement()
	fmt.Println(func1(1, 13))
	fmt.Println(func1(1, 5))
}

func infiniteLoop() {
	for {

	}
}

func forLoop() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println("Sum =>", sum)
}

func whileLoop() {
	sum := 0
	i := 0
	for i < 10 {
		sum += i
		i += 1
	}
	fmt.Println("Sum =>", sum)
}

func ifStatement() {
	sqrt(-2)
	sqrt(9)
}

func sqrt(x float64) {
	if x < 0 {
		fmt.Println("Sqrt =>", math.Sqrt(-x))
	} else {
		fmt.Println("Sqrt =>", math.Sqrt(x))
	}
}

func func1(x, limit int) int {
	if a := (x + 10); a < limit {
		return a
	}
	return x
}
