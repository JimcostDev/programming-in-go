// tienes 3 intentos para adivinar un número aleatorio del 1 al 10, suerte.
package exercises

import (
	"fmt"
	"math/rand"
)

func GuessNumber() {
	var guess int
	number := rand.Intn(10) + 1
	fmt.Println("Adivina el número del 1 al 10")
	for i := 0; i < 3; i++ {
		fmt.Scan(&guess)
		if guess == number {
			fmt.Println("¡Felicidades! Adivinaste el número")
			break
		} else {
			fmt.Println("Número incorrecto")
		}
	}
	fmt.Println("El número era", number)
}
