package concurrency

import (
	"fmt"
	"time"
)

func enviarMensajes(canal chan string) {
	mensajes := []string{"Hola", "Mundo", "Desde", "Goroutines y Canales"}

	for _, mensaje := range mensajes {
		fmt.Printf("Enviando: %s\n", mensaje)
		canal <- mensaje // Enviar mensaje al canal
		time.Sleep(500 * time.Millisecond)
	}

	close(canal) // Cerrar el canal cuando terminamos de enviar datos
}

func Channels() {
	// Crear un canal de tipo string
	canal := make(chan string)

	// Ejecutar la función en una goroutine
	go enviarMensajes(canal)

	// Recibir mensajes desde el canal
	for mensaje := range canal {
		fmt.Printf("Recibido: %s\n", mensaje)
	}

	fmt.Println("Finalizó la comunicación por el canal.")
}
