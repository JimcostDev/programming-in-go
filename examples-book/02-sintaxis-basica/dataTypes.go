package main

import "fmt"

func main() {
	// Declaración de variables con diferentes tipos de datos
	var nombre string = "Hola, mundo!"
	var edad int = 25
	var altura float32 = 1.84
	var esEstudiante bool = true
	var primerLetra rune = 'H'

	// Impresión de los valores
	fmt.Println("Nombre:", nombre)
	fmt.Println("Edad:", edad)
	fmt.Println("Altura:", altura)
	fmt.Println("Es estudiante:", esEstudiante)
	fmt.Println("Primera letra del nombre:", primerLetra)

	// Operaciones básicas
	fmt.Println("La primera letra de tu nombre es un número:", primerLetra)
	fmt.Println("La longitud de tu nombre es:", len(nombre))
}

// Tipos de datos primitivos en Go:

// bool: Representa valores booleanos (true o false).
// string: Cadena de caracteres Unicode.
// int: Entero con signo (tamaño depende de la arquitectura).
// uint: Entero sin signo (tamaño depende de la arquitectura).
// byte: Alias para uint8 (un byte).
// rune: Alias para int32 (un punto de código Unicode).
// float32: Número de punto flotante de precisión simple.
// float64: Número de punto flotante de doble precisión.
