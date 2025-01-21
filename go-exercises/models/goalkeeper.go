package models

import "fmt"

// Struct para un Portero
type Portero struct {
	Nombre string
}

// Implementación del método Jugar para Portero
func (p Portero) Jugar() {
	fmt.Printf("%s está deteniendo los disparos como portero.\n", p.Nombre)
}
