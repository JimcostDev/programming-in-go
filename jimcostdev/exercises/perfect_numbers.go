package exercises

import (
	"fmt"
	"math"
)

// isPrimeNumber verifica si un número es primo de manera eficiente
func isPrimeNumber(n int) bool {
	// Los números menores o iguales a 1 no son primos
	if n <= 1 {
		return false
	}

	// 2 y 3 son números primos
	if n <= 3 {
		return true
	}

	// Si es divisible por 2 o 3, no es primo
	if n%2 == 0 || n%3 == 0 {
		return false
	}

	// Verificamos divisibilidad solo hasta la raíz cuadrada del número
	// Incrementando en pasos de 6 para optimizar (todos los primos > 3 son de la forma 6k±1)
	sqrtN := int(math.Sqrt(float64(n)))
	for i := 5; i <= sqrtN; i += 6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}

	return true
}

// generatePerfectNumbers genera una cantidad específica de números perfectos utilizando primos de Mersenne
func generatePerfectNumbers(cantidad int) []int {
	perfectos := make([]int, 0)
	exponente := 2 // Comenzamos con p=2 ya que es el primer primo

	for len(perfectos) < cantidad {
		// Calculamos el número de Mersenne: 2^p - 1
		mersenne := int(math.Pow(2, float64(exponente))) - 1

		// Si el número de Mersenne es primo
		if isPrimeNumber(mersenne) {
			// Calculamos el número perfecto: 2^(p-1) * (2^p - 1)
			perfecto := int(math.Pow(2, float64(exponente-1))) * mersenne
			perfectos = append(perfectos, perfecto)
		}

		exponente++
	}

	return perfectos
}

// PerfectNumbers es la función principal que ejecuta el programa
func PerfectNumbers() {
	var cantidad int
	fmt.Print("Ingrese la cantidad de números perfectos que desea encontrar: ")
	_, err := fmt.Scan(&cantidad)

	if err != nil {
		fmt.Println("Por favor ingrese un número válido.")
		return
	}

	if cantidad <= 0 {
		fmt.Println("Por favor ingrese un número positivo.")
		return
	}

	fmt.Println("\nCalculando números perfectos...")
	listaPerfectos := generatePerfectNumbers(cantidad)

	fmt.Println("\nNúmeros perfectos encontrados:")
	for i, num := range listaPerfectos {
		fmt.Printf("%d. %d\n", i+1, num)
	}
	fmt.Printf("\nTotal: %d números perfectos.\n", len(listaPerfectos))
}
