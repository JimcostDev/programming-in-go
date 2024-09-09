package exercises

import "fmt"

func PrimeNumbers() {
	// declarar variables
	fmt.Println("Ingrese el número de limitación (máximo que queremos evaluar):")
	var limit int
	fmt.Scan(&limit)

	// mostrar mensaje
	fmt.Println("***** OPCIÓN 1  *****")
	fmt.Printf("Los números primos entre 2 y %v son:\n ", limit)

	// ****** OPCION 1 ****** //
	// logica para saber si es primo
	// cpn =candidatePrimeNumber
	var qtyDivisors int
	for cpn := 2; cpn <= limit; cpn++ {
		qtyDivisors = 0
		for divisor := 1; divisor <= cpn; divisor++ {
			if cpn%divisor == 0 {
				qtyDivisors += 1
			}
		}

		if qtyDivisors == 2 {
			fmt.Println(cpn)
		}
	}

	// probar la función isPrime()

	fmt.Println("***** OPCIÓN 2 *****")
	for x := 2; x <= limit; x++ {
		if isPrime(x) {
			fmt.Println(x)
		}
	}
}

// ****** OPCION 2 ****** //

// isPrime, función para saber si un número es primo
func isPrime(number int) bool {
	if number < 2 {
		return false
	}
	// se excluye el 1 y el mismo numero
	for divisor := 2; divisor < number; divisor++ {
		if number%divisor == 0 {
			return false // no es primo
		}
	}
	return true // es primo
}
