package main

import "fmt"

func Cero(vec [3]int, porc []int) {
	vec[0] = 0
	if len(porc) > 0 {
		porc[0] = 0
	}
	fmt.Println(vec) // [0 2 3]
}

// las porciones son tratadas por referencia, por lo que se pueden modificar en la función
// los vectores son tratados por valor, por lo que no se pueden modificar en la función
func main() {
	v := [3]int{1, 2, 3} // vector
	p := []int{1, 2, 3}  // porción
	Cero(v, p)
	fmt.Println("vector:", v)  // [1 2 3]
	fmt.Println("porción:", p) // [0 2 3]
}
