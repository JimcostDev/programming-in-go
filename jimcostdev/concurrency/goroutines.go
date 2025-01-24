package concurrency

import (
	"fmt"
	"time"
)

// Esta función simula una tarea que toma tiempo
func hacerAlgo() {
	for i := 1; i <= 5; i++ {
		time.Sleep(1 * time.Second) // Simulamos que tarda 1 segundo en hacer algo
		fmt.Println("Tarea #", i)
	}
}

func Goroutines() {
	// Iniciamos la goroutine para ejecutar la función de manera concurrente
	go hacerAlgo()

	// Aquí, el programa principal sigue ejecutándose sin esperar a que termine la goroutine
	fmt.Println("El programa principal sigue ejecutándose...")
	// Esperamos un momento para ver los resultados de la goroutine
	time.Sleep(8 * time.Second)
	fmt.Println("Fin del programa.")
}
