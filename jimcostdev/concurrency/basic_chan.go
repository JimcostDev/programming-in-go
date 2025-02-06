package concurrency

// Usamos make(chan tipo) porque los canales necesitan inicialización dinámica (como slices y maps).
import (
	"fmt"
	"time"
)

func enviarMensaje(canal chan string) {
	time.Sleep(2 * time.Second)          // Simula un retraso
	canal <- "¡Hola desde la Goroutine!" // Enviar mensaje al canal
}

func BasicExampleChannels() {
	// Crear canal de tipo string
	mensajes := make(chan string)

	// Lanzar una Goroutine que enviará un mensaje
	go enviarMensaje(mensajes)

	// Bloquea hasta recibir el mensaje
	fmt.Println("Esperando mensaje...")
	mensaje := <-mensajes // Recibir mensaje
	fmt.Println("Mensaje recibido:", mensaje)
}
