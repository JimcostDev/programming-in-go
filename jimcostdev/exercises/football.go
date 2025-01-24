package exercises

import "fmt"

// Estructura que representa un equipo de fútbol
type Equipo struct {
	nombre string
	goles  int
}

// Función para actualizar el marcador de un equipo
func marcarGol(equipo *Equipo) {
	equipo.goles += 1
	fmt.Printf("¡Gol de %s! Ahora tiene %d goles.\n", equipo.nombre, equipo.goles)
}

// Función para mostrar el marcador actual
func mostrarMarcador(equipo1, equipo2 Equipo) {
	fmt.Printf("\nMarcador actual:\n")
	fmt.Printf("%s: %d\n", equipo1.nombre, equipo1.goles)
	fmt.Printf("%s: %d\n", equipo2.nombre, equipo2.goles)
}

func Football() {
	// Creación de los equipos
	equipo1 := Equipo{nombre: "Barcelona", goles: 0}
	equipo2 := Equipo{nombre: "Real Madrid", goles: 0}

	// Muestra el marcador inicial
	mostrarMarcador(equipo1, equipo2)

	// Marcan algunos goles
	marcarGol(&equipo1) // Gol de Barcelona
	marcarGol(&equipo2) // Gol de Real Madrid
	marcarGol(&equipo1) // Otro gol de Barcelona

	// Muestra el marcador final
	mostrarMarcador(equipo1, equipo2)
}
