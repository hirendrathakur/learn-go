package tour

import "fmt"
import "runtime"
import "time"

// RunSwitches Go only runs the selected case and not all the cases that follows
//Hence switch syntax doesn't require breaks
func RunSwitches() {
	fmt.Println(runtime.GOOS)

	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("Code is running on", os)
	default:
		fmt.Println("Default statement executed")
	}

	say()
}

func say() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good Morning...")
	case t.Hour() < 17:
		fmt.Println("Good Afternoon...")
	default:
		fmt.Println("Good evening...")
	}
}
