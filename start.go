package main

import "fmt"
import "math/rand"
import "math"
import "rsc.io/quote"

func main() {
	fmt.Println(quote.Go())
	fmt.Println("Random number =>", rand.Intn(10))
	fmt.Println("Square Root =>", math.Sqrt(2))
	fmt.Println("Name starting with Capital letter is imported like math.Pi=>", math.Pi)
}
