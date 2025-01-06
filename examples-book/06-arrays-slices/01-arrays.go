package main

import (
	"fmt"
	"reflect"
)

func main() {
	// En Go, los vectores se tratan por valor
	src := [...]int{1, 2, 3, 4, 5} // ... significa que el tamaño del vector se infiere del número de elementos
	var dst [5]int

	// Eso significa que, en asignaciones, se copian
	dst = src
	src[0] = 9

	fmt.Printf("%v != %v \n", src, dst)
	fmt.Println(reflect.TypeOf(dst))
	fmt.Println(reflect.TypeOf(src))
	fmt.Println(reflect.TypeOf(dst) == reflect.TypeOf(src)) // true
}
