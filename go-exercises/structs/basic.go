package structs

import "fmt"

// Definir un struct
type Equipo struct {
	Nombre  string
	Color   string
	Titulos int
}

func Basic() {
	// Crear una instancia de struct
	equipo1 := Equipo{
		Nombre:  "Deportivo Cali",
		Color:   "Verde 💚",
		Titulos: 10,
	}

	// Crear otra instancia usando valores posicionales
	equipo2 := Equipo{"Liverpool", "Rojo ❤️", 19}

	// Acceder y modificar campos
	fmt.Println("Equipo 1:", equipo1)
	fmt.Println("Nombre del Equipo 2:", equipo2.Nombre)

	// Modificar un campo
	equipo1.Titulos = 11
	fmt.Println("Equipo 1 actualizado:", equipo1)

	// Crear un puntero a un struct
	equipo3 := &Equipo{"Real Madrid", "Blanco 🤍", 35}
	equipo3.Titulos++
	fmt.Println("Equipo 3 (con puntero):", *equipo3)
}
