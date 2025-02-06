# 📌 Fundamentos de los Canales en Go

Los **canales** en Go permiten la comunicación segura entre **Goroutines**, evitando condiciones de carrera sin necesidad de sincronización manual con `sync.Mutex` o `sync.WaitGroup`.

## 📌 Conceptos Claves
✅ Se crean con `make(chan tipo)`, no con `var`, porque son estructuras dinámicas de memoria.  
✅ Se usa `<-` para **enviar** y **recibir** datos a través del canal.  
✅ Son **bloqueantes** por defecto: la Goroutine que **envía** espera a que alguien **reciba**, y viceversa.  

## 🔹 Ejemplo Práctico: Canal de Mensajes
El siguiente código simula una Goroutine que envía un mensaje y otra que lo recibe:

```go
package main

import (
	"fmt"
	"time"
)

func enviarMensaje(canal chan string) {
	time.Sleep(2 * time.Second) // Simula un retraso
	canal <- "¡Hola desde la Goroutine!" // Enviar mensaje al canal
}

func main() {
	// Crear canal de tipo string
	mensajes := make(chan string)

	// Lanzar una Goroutine que enviará un mensaje
	go enviarMensaje(mensajes)

	// Bloquea hasta recibir el mensaje
	fmt.Println("Esperando mensaje...")
	mensaje := <-mensajes // Recibir mensaje
	fmt.Println("Mensaje recibido:", mensaje)
}
```

## 📌 Explicación
1️⃣ **Crear el canal:**  
```go
mensajes := make(chan string)
```
Se usa `make(chan tipo)` porque los canales necesitan inicialización dinámica.

2️⃣ **Enviar un dato al canal:**  
```go
canal <- "¡Hola desde la Goroutine!"
```
📌 **El símbolo `<-` aquí envía el mensaje** al canal.

3️⃣ **Recibir un dato del canal:**  
```go
mensaje := <-mensajes
```
📌 **Aquí `<-` recibe el mensaje** desde el canal.

4️⃣ **Bloqueo y sincronización:**  
   - **La Goroutine espera 2 segundos antes de enviar el mensaje.**  
   - **`main()` queda bloqueado en `mensaje := <-mensajes` hasta que recibe el dato.**

## 🔹 Uso de `<-` en Canales
| Operación             | Uso                         | Explicación |
|----------------------|-------------------------|-------------|
| **Escribir al canal** | `canal <- valor`        | Envía un valor al canal |
| **Leer del canal**    | `variable := <-canal`   | Recibe un valor del canal |

---

### 🚀 **Conclusión**
Los canales en Go facilitan la comunicación segura entre **Goroutines**, permitiendo un modelo de concurrencia basado en el paso de mensajes en lugar de compartir memoria. ¡Son una herramienta poderosa para la programación concurrente! 🔥
