package exercises

import (
	"fmt"
)

/*
   ---- CONJETURA DE GOLDBACH ----
   Todo número par mayor que 2 se puede expresar como suma de dos números primos.

   Para el número 14 tendríamos: 3 y 11, 7 y 7
   (1 y 13 no valdrian, ya que el uno no es considerado un primo)

   Hacer un programa que muestre todas las parejas de primos en las que se puede
   expresar un número par mayor que 2.
*/

func Goldbach() {
	//declarar variables
	var numero int
	fmt.Println("Digite un número par mayor que 2: ")

	//capturar numero
	fmt.Scanln(&numero)

	//condición para verificar que el numero sea par y mayor que 2
	if numero%2 == 0 && numero > 2 {
		encontrado := false
		for a := 2; a < numero; a++ {
			//se toma num(14) y se le resta el primer numero primo a(3) y si el resultado b(11) es un numero primo, se forma una pareja (a,b) y asi hasta terminar el ciclo.
			if EsPrimo(a) {
				b := numero - a
				if EsPrimo(b) {
					encontrado = true // es primo b
					//condición para evitar que se repitan las parejas
					if a <= b {
						fmt.Println("Primos: ", a, " ", b)
					}
				}
			}
		}
		if !encontrado {
			fmt.Println("No se ha encontrado ninguna pareja para el número: ", numero)
		}
	} else {
		fmt.Println("No es un numero valido")
	}
}

// EsPrimo funcion para saber si un número es primo
func EsPrimo(n int) bool {
	if n < 2 {
		return false
	}
	//se excluye el  1 y el mismo numero
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false // no es primo
		}
	}
	return true //es primo
}
