package unittests

// La función Factorial retorna el factorial de un número
func Factorial(n uint64) uint64 {
	if n == 0 {
		return 1
	}
	return n * Factorial(n-1)
}

// La función Sumar suma los elementos de una lista de enteros
func SumarEnteros(lista []int) int {
	suma := 0
	for _, elemento := range lista {
		suma += elemento
	}
	return suma
}
