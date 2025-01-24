package main

import (
	"fmt"

	_ "github.com/JimcostDev/programming-in-go/go-exercises/arrays_slices"
	_ "github.com/JimcostDev/programming-in-go/go-exercises/defer_panic"
	"github.com/JimcostDev/programming-in-go/go-exercises/exercises"
	"github.com/JimcostDev/programming-in-go/go-exercises/structs"
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
	type edad uint8 // 0 a 255
	var edadRonaldo edad = 25
	fmt.Printf("La edad de %s es %d\n", exercises.Nombre, edadRonaldo)

	// tabla := exercises.Iterar()
	// fmt.Println(tabla)

	//files.ReadFile()
	structs.Basic()
	exercises.Pointers()
}
