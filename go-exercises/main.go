package main

import (
	"fmt"

	"github.com/JimcostDev/programming-in-go/go-exercises/arrays_slices"
	"github.com/JimcostDev/programming-in-go/go-exercises/exercises"
	//"github.com/JimcostDev/programming-in-go/go-exercises/files"
)

func main() {
	// aqui se llama a la función que se desea ejecutar, es decir el ejercicio que se desea probar
	exercises.Nombre = "Ronaldo Jiménez" // Nombre se toma de vars.go
	// función anonima para saludar
	saludar := func(name string) string {
		return "Hola " + name
	}
	fmt.Println(saludar(exercises.Nombre))

	// tabla := exercises.Iterar()
	// fmt.Println(tabla)

	//files.ReadFile()
	arrays_slices.Equipos()
}
