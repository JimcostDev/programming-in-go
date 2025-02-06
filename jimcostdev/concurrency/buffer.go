package concurrency

// El uso de buffers en los canales permite enviar y recibir mensajes sin bloqueo.
// Son como una sala de espera o un almacén temporal que permite almacenar
// varios mensajes sin que el remitente tenga que esperar a que alguien los reciba inmediatamente.
import (
	"fmt"
)

func Buffer() {
	// Canal con buffer de 2
	mensajes := make(chan string, 2)

	// Podemos enviar 2 mensajes sin bloqueo
	mensajes <- "Hola"
	mensajes <- "Mundo"
	//mensajes <- "buffer " // ¡SE BLOQUEA! porque buffer está lleno

	// Leemos los mensajes cuando queramos
	fmt.Println(<-mensajes)
	fmt.Println(<-mensajes)
	// fmt.Println(<-mensajes) // all goroutines are asleep - deadlock!
}
