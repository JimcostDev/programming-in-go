package exercises

import (
	"fmt"
)

func Pow() {
	//simular un ejercicio de potenciacón matematica
	var base int
	var exponente int
	resultado := 1 //acumulador

	fmt.Println("Ingrese la base: ")
	fmt.Scan(&base)

	fmt.Println("Ingrese el exponente: ")
	fmt.Scan(&exponente)

	for i := 0; i < exponente; i++ {
		resultado *= base //resultado = resultado * base
	}

	fmt.Println("El resultado es: ", resultado)
}
