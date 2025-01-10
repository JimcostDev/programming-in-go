package main

import "fmt"

func Suma(n ...int) int {
	total := 0
	for _, v := range n {
		total += v
	}
	return total
}

func main() {
	fmt.Println(Suma(1, 2, 3))       // 6
	fmt.Println(Suma(1, 2, 3, 4, 5)) // 15
	fmt.Println(Suma())              // 0

	valores := []int{7, 6, 8, 4}
	fmt.Println(Suma(valores...)) // 25
}
