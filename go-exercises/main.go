package main

import (
	"fmt"

	"github.com/JimcostDev/programming-in-go/go-exercises/exercises"
)

func main() {
	// aqui se llama a la función que se desea ejecutar, es decir el ejercicio que se desea probar
	exercises.Nombre = "Ronaldo Jiménez"
	fmt.Println(exercises.Nombre)

	// convertir número a texto
	estado, texto := exercises.ConvertirNumeroTexto(105)
	fmt.Println(estado, texto)
}
