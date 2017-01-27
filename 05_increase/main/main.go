package main

import (
	"fmt"
)

func wrapper() func() int {
	x :=0
	println(x)
	return func() int {
		x++
		println(x)
		return x
	}

}
func main() {
	increment := wrapper()
	fmt.Println(increment())
	fmt.Println(increment())
}
