package main

import "fmt"
import "time"

func main() {
	i := 2
	fmt.Print("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("it's a weekday")
	}

	t := time.Now()
	switch {
	case t.Hour() <12:
		fmt.Println("it's before noon")
	default:
		fmt.Println("it's after noon")
	}
	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I am a bool.")
		case int :
			fmt.Println("I am an int.")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(3)
	whatAmI("OH")
}
