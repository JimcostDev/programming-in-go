package exercises

import (
	"fmt"
)

// Función para obtener un número del usuario
func obtenerNumero(orden string) int {
	var numero int
	fmt.Printf("Ingrese el %s número: ", orden)
	fmt.Scan(&numero)
	return numero
}

// Función para encontrar el número mayor
func numeroMayor(numeros ...int) int {
	mayor := numeros[0]
	for _, num := range numeros { // _ ignorar indice
		if num > mayor {
			mayor = num
		}
	}
	return mayor
}

func LargestNumber() {
	primero := obtenerNumero("primer")
	segundo := obtenerNumero("segundo")
	tercero := obtenerNumero("tercer")
	cuarto := obtenerNumero("cuarto")

	mayor := numeroMayor(primero, segundo, tercero, cuarto)
	fmt.Printf("El número mayor es: %d\n", mayor)
}
