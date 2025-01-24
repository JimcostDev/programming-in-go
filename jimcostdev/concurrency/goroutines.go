package concurrency

import (
	"fmt"
	"strings"
	"time"
)

// Goroutines : Función que muestra el uso de goroutines
func Goroutines(name string) {
	// goroutine es una función que se ejecuta de manera concurrente con otras funciones
	// la palabra clave go permite ejecutar una función como goroutine
	letters := strings.Split(name, "")
	for _, letter := range letters {
		// mostar cada 1s una letra de la palabra
		time.Sleep(500 * time.Millisecond)
		fmt.Println(letter)
	}
}
