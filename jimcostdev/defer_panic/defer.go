package defer_panic

import "fmt"

/*
defer se utiliza para posponer la ejecución de una función o una instrucción
hasta que la función que la contiene esté a punto de finalizar. Es especialmente
útil para manejar recursos que necesitan ser liberados, como cerrar archivos,
conexiones de base de datos, o liberar memoria.
*/

func Defer() {
	fmt.Println("Inicio del programa")

	defer fmt.Println("Esto se ejecuta al final") // Pospone esta instrucción hasta antes de que main termine
	fmt.Println("Ejecutando lógica principal")

	fmt.Println("Fin de la lógica principal")
}
