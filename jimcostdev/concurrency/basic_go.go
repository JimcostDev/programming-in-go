package concurrency

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func imprimirMensaje(mensaje string) {
	defer wg.Done() // Indica que la Goroutine ha finalizado
	for i := 0; i < 5; i++ {
		fmt.Println(mensaje, i)
		time.Sleep(500 * time.Millisecond)
	}
}

func BasicExampleGoroutines() {
	wg.Add(1)
	go imprimirMensaje("Hola desde la Goroutine")

	wg.Add(1)
	go imprimirMensaje("Hola desde otra Goroutine")

	wg.Wait() // Espera a que todas las Goroutines terminen
	fmt.Println("Finalizando el programa")
}
