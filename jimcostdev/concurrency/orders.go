package concurrency

import (
	"fmt"
	"time"
)

func procesarPedido(pedidos chan int, resultados chan string) {
	for id := range pedidos { // Leemos continuamente del canal de pedidos
		// Simulamos procesamiento
		time.Sleep(time.Second)

		// Enviamos el resultado
		resultados <- fmt.Sprintf("Pedido %d procesado", id)
	}
}

func Orders() {
	pedidos := make(chan int, 3)
	resultados := make(chan string, 3)

	// Iniciamos el procesador de pedidos
	go procesarPedido(pedidos, resultados)

	// Enviamos 3 pedidos
	for i := 1; i <= 3; i++ {
		pedidos <- i
		fmt.Printf("Pedido %d enviado\n", i)
	}

	// Recibimos los resultados
	for i := 1; i <= 3; i++ {
		resultado := <-resultados
		fmt.Println(resultado)
	}

	// Cerramos el canal de pedidos
	close(pedidos)
}
