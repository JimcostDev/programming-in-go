package main

import "fmt"

func main() {
	var nombre string = "Ronaldo"
	var edad int = 23
	var altura float64 = 1.90
	var esEstudiante bool = true
	var caracter rune = 'R'

	fmt.Println("Ejemplo de uso de auxiliares de formato:")
	fmt.Printf("Nombre: %s\n", nombre)
	fmt.Printf("Nombre: %v\n", nombre)
	fmt.Printf("Edad: %d años\n", edad)
	fmt.Printf("Altura: %.2f metros\n", altura)
	fmt.Printf("Es estudiante: %t\n", esEstudiante)
	fmt.Printf("Primer letra del nombre: %c\n", caracter)
	fmt.Printf("Edad en binario: %b\n", edad)
	fmt.Printf("Altura en notación científica: %e\n", altura)
	fmt.Printf("Edad en hexadecimal: %x\n", edad)
	fmt.Printf("Tipo de la variable 'nombre': %T\n", nombre)

	// Ejemplo con modificadores
	fmt.Printf("%-10s %05d %.2f\n", nombre, edad, altura) // Alinea a la izquierda, rellena con ceros, precisión de 2 decimales
}

// Se utiliza Printf para mostrar los valores de las variables utilizando los siguientes auxiliares:
// %s: Para cadenas de texto.
// %d: Para enteros decimales.
// %f: Para números de punto flotante.
// %t: Para valores booleanos.
// %c: Para caracteres individuales.
// %b: Para representación binaria de un entero.
// %e: Para notación científica.
// %x: Para representación hexadecimal.
// %T: Para mostrar el tipo de una variable.
// %v: Para imprimir el valor de una variable en su representación más legible.
