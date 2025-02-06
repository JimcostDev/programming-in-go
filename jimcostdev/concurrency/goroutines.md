# 📌 Fundamentos de las Goroutines en Go

Las **Goroutines** son funciones o métodos que se ejecutan de manera concurrente en Go. Son ligeras y permiten ejecutar múltiples tareas simultáneamente.

## 📌 Conceptos Claves
✅ Se ejecutan con la palabra clave `go`.  
✅ Son más ligeras que los hilos del sistema operativo.  
✅ Se ejecutan en **Go Runtime**, que gestiona su planificación.

## 🔹 Ejemplo Práctico: Uso de Goroutines
El siguiente código muestra cómo una Goroutine ejecuta una función de manera concurrente:

```go
package main

import (
	"fmt"
	"time"
)

func imprimirMensaje(mensaje string) {
	for i := 0; i < 5; i++ {
		fmt.Println(mensaje, i)
		time.Sleep(500 * time.Millisecond) // Simula un proceso
	}
}

func main() {
	// Ejecutar función como una Goroutine
	go imprimirMensaje("Hola desde la Goroutine")

	// Ejecutar otra función en el hilo principal
	imprimirMensaje("Hola desde main")
}
```

## 📌 Explicación
1️⃣ **Lanzar una Goroutine:**  
```go
go imprimirMensaje("Hola desde la Goroutine")
```
📌 La función `imprimirMensaje` se ejecutará concurrentemente con `main()`.

2️⃣ **Ejecutar en el hilo principal:**  
```go
imprimirMensaje("Hola desde main")
```
📌 Esta función se ejecuta de manera normal en el hilo principal.

3️⃣ **Problema con la ejecución:**  
El programa puede finalizar antes de que la Goroutine termine, porque `main()` no espera su ejecución.

## 🔹 Solución: Usar `sync.WaitGroup`
Podemos utilizar `sync.WaitGroup` para asegurarnos de que `main()` espere a que terminen las Goroutines.

```go
package main

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

func main() {
	wg.Add(1)
	go imprimirMensaje("Hola desde la Goroutine")

	wg.Add(1)
	go imprimirMensaje("Hola desde otra Goroutine")

	wg.Wait() // Espera a que todas las Goroutines terminen
	fmt.Println("Finalizando el programa")
}
```

## 🔹 Ventajas de las Goroutines
| Ventaja           | Explicación |
|------------------|-------------|
| **Ligeras** | Ocupan menos memoria que los threads del sistema operativo. |
| **Fáciles de usar** | Solo se necesita `go` para lanzarlas. |
| **Manejo eficiente** | Go Runtime gestiona su planificación y uso de CPU. |

---

### 🚀 **Conclusión**
Las Goroutines permiten la concurrencia en Go de manera eficiente y sencilla. Sin embargo, es importante sincronizar su ejecución correctamente para evitar que el programa termine antes de que las Goroutines finalicen. 🔥
