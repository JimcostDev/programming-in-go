/*
En Go, un receptor es un valor que aparece en la definición de un método.
Este valor, que suele ser una instancia de una estructura, se pasa
implícitamente al método cuando se invoca. Es decir, los métodos en Go

	están asociados a tipos específicos.

Un método es simplemente una función con un receptor.
Esto permite que los métodos accedan y modifiquen los campos
del receptor, creando un concepto similar a los métodos de los
objetos en otros lenguajes orientados a objetos.
*/
package structs

import "fmt"

// Definimos una estructura para representar un círculo
type Circle struct {
	radius float64
}

// Método para calcular el área del círculo
func (c Circle) area() float64 {
	return 3.14 * c.radius * c.radius
}

func FunctionReceptors() {
	circle1 := Circle{radius: 5}
	fmt.Println("Área del círculo:", circle1.area())
}
