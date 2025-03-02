package exercises

// El juego de "Picas y Fijas" consiste en adivinar un número
// secreto compuesto por dígitos únicos. Primero, selecciona
// la cantidad de cifras (entre 2 y 9). Luego, intenta adivinar
// el número secreto ingresando tu intento.

// Cada intento te proporcionará dos pistas: Fijas indican
// los dígitos que están en la posición correcta y Picas
// los dígitos correctos pero en posiciones distintas.
// Usa estas pistas para ir deduciendo el número correcto.
// Ejemplo:
// 	- Número secreto: 1234
// 	- Intento: 4321
// 	- Picas: 4
// 	- Fijas: 0

import (
	"fmt"
	"strconv"

	"github.com/JimcostDev/jimcostdev-utils/random"
)

// solicitarCantidadCifras pide al usuario que elija la cantidad de cifras para el juego
func solicitarCantidadCifras() int {
	var digits int

	for {
		fmt.Print("Selecciona la cantidad de cifras (entre 2 y 9): ")
		_, err := fmt.Scan(&digits)

		if err == nil && digits >= 2 && digits <= 9 {
			break
		}

		fmt.Println("Por favor, ingresa un número válido entre 2 y 9.")
	}

	return digits
}

// generarNumeroSecreto crea un número aleatorio con dígitos únicos
func generarNumeroSecreto(digits int) string {
	numeroSecreto := random.UniqueDigits(digits)
	return numeroSecreto
}

// jugarPicasFijas maneja la lógica principal del juego
func jugarPicasFijas(secretNumber string) {
	fmt.Printf("Se ha generado un número secreto de %d cifras. ¡Adivínalo!\n", len(secretNumber))

	intentos := 0
	adivinado := false

	for !adivinado {
		intentos++

		// Solicitar intento al usuario
		intento := solicitarIntento(len(secretNumber))

		// Calcular picas y fijas
		picas, fijas := calcularPicasFijas(secretNumber, intento)

		// Mostrar resultado
		fmt.Printf("Intento #%d: %s -> Picas: %d, Fijas: %d\n", intentos, intento, picas, fijas)

		// Verificar si adivinó
		if fijas == len(secretNumber) {
			adivinado = true
			fmt.Printf("¡Felicidades! Has adivinado el número secreto %s en %d intentos.\n", secretNumber, intentos)
		}
	}
}

// solicitarIntento pide al usuario que ingrese su intento
func solicitarIntento(longitud int) string {
	var intento string

	for {
		fmt.Printf("Ingresa tu intento (un número de %d cifras con dígitos únicos): ", longitud)
		fmt.Scan(&intento)

		if validarIntento(intento, longitud) {
			break
		}

		fmt.Println("Intento inválido. Asegúrate de ingresar la cantidad correcta de cifras y que los dígitos no se repitan.")
	}

	return intento
}

// validarIntento verifica que el intento cumpla con las reglas
func validarIntento(intento string, longitud int) bool {
	// Verificar longitud
	if len(intento) != longitud {
		return false
	}

	// Verificar que sean dígitos
	_, err := strconv.Atoi(intento)
	if err != nil {
		return false
	}

	// Verificar que no haya dígitos repetidos
	digitSet := make(map[rune]bool)
	for _, digit := range intento {
		if digitSet[digit] {
			return false
		}
		digitSet[digit] = true
	}

	return true
}

// calcularPicasFijas determina la cantidad de picas y fijas
func calcularPicasFijas(secretNumber, intento string) (int, int) {
	picas := 0
	fijas := 0

	// Verificar fijas (dígitos correctos en posición correcta)
	for i := 0; i < len(secretNumber); i++ {
		if secretNumber[i] == intento[i] {
			fijas++
		}
	}

	// Conjunto de dígitos que aparecen en ambos números
	digitosComunes := 0
	secretDigits := make(map[rune]int)
	guessDigits := make(map[rune]int)

	// Contar ocurrencias de cada dígito
	for _, digit := range secretNumber {
		secretDigits[digit]++
	}

	for _, digit := range intento {
		guessDigits[digit]++
	}

	// Contar dígitos comunes
	for digit, count := range secretDigits {
		if guessCount, exists := guessDigits[digit]; exists {
			digitosComunes += min(count, guessCount)
		}
	}

	// Las picas son los dígitos comunes menos las fijas
	picas = digitosComunes - fijas

	return picas, fijas
}

// PicasFijas ejecuta el juego principal de Picas y Fijas
func PicasFijas() {
	fmt.Println("¡Bienvenido al juego de Picas y Fijas!")

	// Solicitar cantidad de cifras
	digits := solicitarCantidadCifras()

	// Generar número secreto
	secretNumber := generarNumeroSecreto(digits)

	// Iniciar el juego
	jugarPicasFijas(secretNumber)
}
