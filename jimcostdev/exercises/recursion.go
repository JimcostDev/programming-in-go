package exercises

import (
	"fmt"
)

// Recursion is a function that calls itself
func Exponentiation(value int) {
	if value > 1024 {
		fmt.Println("Value is too high")
		return
	}
	fmt.Println(value)
	Exponentiation(value * 2)
}
