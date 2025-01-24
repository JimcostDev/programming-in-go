package structs

import "github.com/JimcostDev/programming-in-go/jimcostdev/models"

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
	enganche := models.Enganche{
		Delantero:   models.Delantero{Nombre: "James Rodríguez", Goles: 15},
		Asistencias: 12,
	}

	// Usar la interfaz en un partido
	partido(portero)
	partido(delantero)
	partido(enganche)
}
