package middleware

import "fmt"

// Operaciones básicas
func sum(a, b int) int {
	return a + b
}

func sub(a, b int) int {
	return a - b
}

func mult(a, b int) int {
	return a * b
}

// Middleware que toma una función de operación y la modifica (puedes agregar log, validación, etc)
func operacionesMidd(operacion func(int, int) int) func(int, int) int {
	return func(a, b int) int {
		fmt.Println("Ejecutando operación en middleware")
		// Aquí puedes modificar el flujo, como hacer logging o validaciones
		result := operacion(a, b)
		// También podrías realizar alguna modificación en el resultado antes de devolverlo
		return result
	}
}

// Función que usa el middleware para envolver las operaciones
func MyMiddleware() {
	fmt.Println("Inicio")
	// Aplicando middleware a la operación de suma
	result := operacionesMidd(sum)(5, 3)
	fmt.Println("Resultado de la operación: ", result)

	// Aplicando middleware a la operación de resta
	result = operacionesMidd(sub)(5, 3)
	fmt.Println("Resultado de la operación: ", result)

	// Aplicando middleware a la operación de multiplicación
	result = operacionesMidd(mult)(5, 3)
	fmt.Println("Resultado de la operación: ", result)
}
