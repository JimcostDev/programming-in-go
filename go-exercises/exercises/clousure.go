package exercises

import "fmt"

// Closure: una función que devuelve otra función
func contador() func() int {
	count := 0 // Variable local que se mantiene en el closure
	return func() int {
		count++ // Incrementa la variable count cada vez que se llama
		return count
	}
}

// Se pueden usar closures para generar funciones que filtren datos según ciertos criterios.
func filtroMayorQue(n int) func(int) bool {
	return func(x int) bool {
		return x > n
	}
}

func Clouseres() {
	// Creamos un closure con la función contador
	contar := contador()

	// Llamamos al closure varias veces
	fmt.Println(contar()) // Imprime: 1
	fmt.Println(contar()) // Imprime: 2
	fmt.Println(contar()) // Imprime: 3

	// Creamos otro closure independiente
	otroContar := contador()
	fmt.Println(otroContar()) // Imprime: 1
	fmt.Println()

	// Creamos un closure con la función filtroMayorQue
	esMayorQue10 := filtroMayorQue(10)
	fmt.Println(esMayorQue10(15)) // true
	fmt.Println(esMayorQue10(5))  // false

}
