package concurrency

// El uso de buffers en los canales permite enviar y recibir mensajes sin bloqueo.
// Son como una sala de espera o un almacén temporal que permite almacenar
// varios mensajes sin que el remitente tenga que esperar a que alguien los reciba inmediatamente.
import (
	"fmt"
)

// problema: deadlock
// solucion 1: aumentar el buffer a 3
// solucion 2: liberar un espacio en el buffer
func Buffer() {
	// Canal con buffer de 2
	mensajes := make(chan string, 2) // aumentar a 3 (solucion 1)

	// Podemos enviar 2 mensajes sin bloqueo
	mensajes <- "Hola"  // ✅ Se almacena en el buffer
	mensajes <- "Mundo" // ✅ Se almacena en el buffer
	// fmt.Println(<-mensajes) // Se libera un espacio en el buffer (solucion 2)
	mensajes <- "buffer" // ❌ Aquí ocurre el problema, el buffer está lleno - deadlock.
	close(mensajes)

	// Leemos los mensajes del canal
	for ms := range mensajes {
		fmt.Println("Recibiendo: ", ms)
	}
	fmt.Println("Canal cerrado. Fin!")
}
