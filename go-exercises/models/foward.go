package models

import "fmt"

// Struct para un Delantero
type Delantero struct {
	Nombre string
	Goles  int
}

// Implementación del método Jugar para Delantero
func (d Delantero) Jugar() {
	fmt.Printf("%s está marcando goles como delantero. Lleva %d goles en la temporada.\n", d.Nombre, d.Goles)
}
