package models

import "fmt"

type Enganche struct {
	Delantero
	Asistencias int
}

// Sobrescribir o extender el método Jugar
func (e Enganche) Jugar() {
	fmt.Printf("%s está organizando el juego como enganche. Lleva %d asistencias y %d goles en la temporada.\n", e.Nombre, e.Asistencias, e.Goles)
}
