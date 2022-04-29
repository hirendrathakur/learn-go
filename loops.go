package main

import "fmt"
import "math"

func main() {
	for_loop()
	while_loop()
	if_statement()
	fmt.Println(func_1(1, 13))
	fmt.Println(func_1(1, 5))
}

func infinite_loop() {
	for {

	}
}

func for_loop() {
	sum := 0
	for i:=0; i < 10; i++ {
		sum+=i
	}
	fmt.Println("Sum =>", sum)	
}

func while_loop() {
	sum := 0
	i := 0
	for i<10 {
		sum+=i
		i+=1
	}
	fmt.Println("Sum =>", sum)
}

func if_statement() {
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

func func_1(x, limit int) int {
	if a := (x+10); a < limit {
		return a
	} 
	return x
}