package concurrency

/*
sync.WaitGroup es una estructura que nos permite esperar a que un grupo de Goroutines termine su ejecución
¿Cuándo Usar sync.WaitGroup?
- Cuando necesitas esperar a que todas las Goroutines terminen antes de continuar.
- Para ejecutar tareas en paralelo sin que el programa principal termine prematuramente.
- Para mejorar la eficiencia en programas concurrentes.
*/

import (
	"fmt"
	"sync"
	"time"
)

func tarea(id int, wg *sync.WaitGroup) {
	defer wg.Done() // Indica que esta Goroutine ha terminado

	fmt.Printf("Tarea %d iniciando...\n", id)
	time.Sleep(time.Second * 2) // Simula una tarea que toma tiempo
	fmt.Printf("Tarea %d finalizada\n", id)
}

func SyncWG() {
	var wg sync.WaitGroup // Declaramos una WaitGroup

	numTareas := 5 // Número de tareas a ejecutar en paralelo

	for i := 1; i <= numTareas; i++ {
		wg.Add(1)        // Añadimos una tarea al WaitGroup
		go tarea(i, &wg) // Lanzamos la Goroutine
	}

	wg.Wait() // Esperamos a que todas las Goroutines terminen
	fmt.Println("Todas las tareas han finalizado")
}
