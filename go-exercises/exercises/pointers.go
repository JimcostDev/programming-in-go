package exercises

import "fmt"

func Pointers() {
	// Una caja fuerte con 100 euros
	dinero := 100

	// Un guardia con la combinación de la caja fuerte
	guardia := &dinero
	fmt.Println("Hay", dinero, "euros en la caja fuerte.")
	fmt.Println("El guardia sabe que hay", *guardia, "euros en la caja fuerte.")

	// dirección de memoria de la variable dinero desde dinero y guardia
	fmt.Println("La dirección de memoria, dinero: ", &dinero)
	fmt.Println("La dirección de memoria guardia: ", guardia)

	// El guardia cambia el dinero a 200 euros
	*guardia = 200

	fmt.Println("Ahora hay", dinero, "euros en la caja fuerte.")
	fmt.Println("")

	// otro ejemplo
	fmt.Println("Ejemplo de comparación de direcciones de memoria")
	i := 0
	x := 0

	y := &i
	z := &x

	fmt.Println("La dirección de memoria de i es: ", y)
	fmt.Println("La dirección de memoria de x es: ", z)
	if y == z {

		fmt.Println("son iguales, tienen el mismo valor")

	} else {

		fmt.Println("son diferentes, apuntan a diferente dirección de memoria")

	}
}
