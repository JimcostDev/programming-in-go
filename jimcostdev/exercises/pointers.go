package exercises

// Los apuntadores (o punteros) permiten almacenar la dirección de memoria de una variable.
// En Go, los apuntadores se representan con el símbolo * seguido del tipo de dato al que apunta.
// Para obtener la dirección de memoria de una variable, se utiliza el símbolo & seguido del nombre de la variable.
// Solo pueden apuntar a un tipo en concreto, no se pueden mezclar tipos de datos.
import "fmt"

func Pointers() {
	exampleOne()
	exampleTwo()
	exampleThree()
}

func exampleOne() {
	// declarando un puntero de ints
	var puntero *int

	variable := 10
	puntero = &variable // obtengo la dirección de memoria de la variable

	fmt.Printf("El valor de la variable es: %v\n", variable)
	fmt.Printf("La dirección de memoria de la variable es: %v\n", &variable)
	fmt.Printf("El valor del puntero es: %v\n", *puntero)
	fmt.Printf("La dirección de memoria del puntero es: %v\n", puntero)

	// Cambiando el valor de la variable a través del puntero
	*puntero = 20
	fmt.Printf("El valor nuevo es: %v\n", variable)
}

func exampleTwo() {
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

func exampleThree() {
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
}
