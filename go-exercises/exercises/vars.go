package exercises

import (
	"fmt"
	"time"
)

var Nombre string // variable global, se puede acceder desde cualquier parte del paquete por ser pública (mayúscula inicial)
var Estado bool
var Sueldo float32
var Fecha time.Time

func Vars() {
	// enteros
	ShowIntegers()

	Nombre = "Ronaldo"
	Estado = true
	Sueldo = 1876.33
	Fecha = time.Now()

	// mostrar variables
	fmt.Println(Nombre)
	fmt.Println(Estado)
	fmt.Println(Sueldo)
	fmt.Println(Fecha)
}
