package exercises

// La implementación usa varias optimizaciones:

// Solo verifica hasta la raíz cuadrada del número
// Elimina verificaciones innecesarias de números pares
// Usa el patrón 6k ± 1 para reducir el número de divisiones
// Implementa la Criba de Eratóstenes para generar múltiples primos eficientemente

import (
	"fmt"
	"math"
)

// IsPrime verifica si un número es primo usando optimizaciones
func IsPrime(n int) bool {
	// Casos especiales
	if n <= 1 {
		return false
	}
	if n <= 3 {
		return true
	}
	if n%2 == 0 || n%3 == 0 {
		return false
	}

	// Solo necesitamos verificar hasta la raíz cuadrada del número
	// Y solo necesitamos verificar los números que son de la forma 6k ± 1
	limit := int(math.Sqrt(float64(n)))
	for i := 5; i <= limit; i += 6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}
	return true
}

// GeneratePrimes genera una lista de números primos hasta n usando la Criba de Eratóstenes
func GeneratePrimes(n int) []int {
	// Crear slice para la criba
	isPrime := make([]bool, n+1)
	for i := 2; i <= n; i++ {
		isPrime[i] = true
	}

	// Aplicar la criba
	for i := 2; i*i <= n; i++ {
		if isPrime[i] {
			for j := i * i; j <= n; j += i {
				isPrime[j] = false
			}
		}
	}

	// Recolectar los números primos
	primes := make([]int, 0)
	for i := 2; i <= n; i++ {
		if isPrime[i] {
			primes = append(primes, i)
		}
	}
	return primes
}

func PrimesImprovedVersion() {
	// Verificar si un número es primo
	numero := 97
	if IsPrime(numero) {
		fmt.Printf("%d es primo\n", numero)
	}

	// Generar primos hasta 100
	primos := GeneratePrimes(100)
	fmt.Println("Primos hasta 100:", primos)
}
