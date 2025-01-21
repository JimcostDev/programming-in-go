package structs

import "github.com/JimcostDev/programming-in-go/go-exercises/models"

// Definir la interfaz
type Jugador interface {
	Jugar()
}

// Función que usa la interfaz
func partido(j Jugador) {
	j.Jugar()
}

func Interfaces() {
	// Crear jugadores
	portero := models.Portero{Nombre: "Jan Oblak"}
	delantero := models.Delantero{Nombre: "Robert Lewandowski", Goles: 28}

	// Usar la interfaz en un partido
	partido(portero)
	partido(delantero)
}
