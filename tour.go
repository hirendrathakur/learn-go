package main

import (
	"com.learn/learn-go/tour"
	"fmt"
	"time"
)

func saySomething(s string, ch chan int) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
	ch <- 0
}

func main() {
	RunStart()
	tour.RunDataTypes()
	tour.RunFunctions()
	tour.RunLoops()
	tour.RunMethods()
	tour.RunPointers()
	tour.RunSwitches()
	tour.RunVariables()
}

func RunStart() {
	ch := make(chan int)
	go saySomething("Hi", ch)
	fmt.Println(<-ch)
	fmt.Println("Done")
}
