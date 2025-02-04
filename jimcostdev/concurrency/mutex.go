package concurrency

/*
 ¿Cuándo Usar sync.Mutex?
🔹 Cuando varias Goroutines acceden/modifican la misma variable.
🔹 Para evitar condiciones de carrera en datos compartidos.
🔹 Para proteger secciones críticas en concurrencia.
sync.Mutex es una herramienta clave en Go para mantener la seguridad en la concurrencia.
*/

import (
	"fmt"
	"sync"
	"time"
)

var contador int = 0 // Variable compartida
var mutex sync.Mutex // Declaración de un Mutex

func incrementar(id int, wg *sync.WaitGroup) {
	defer wg.Done() // Indica que la Goroutine ha finalizado

	// Bloquea la sección crítica
	mutex.Lock()
	contador++
	fmt.Printf("Goroutine %d incrementó el contador a %d\n", id, contador)
	// Desbloquea la sección crítica
	mutex.Unlock()

	time.Sleep(time.Millisecond * 500) // Simula una tarea
}

func Mutex() {
	var wg sync.WaitGroup
	numGoroutines := 5

	for i := 1; i <= numGoroutines; i++ {
		wg.Add(1)
		go incrementar(i, &wg)
	}

	wg.Wait() // Espera a que todas las Goroutines terminen
	fmt.Println("Valor final del contador:", contador)
}
