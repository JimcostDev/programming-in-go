package files

// importar paquetes de ejercicios
import (
	"fmt"
	"io"
	"os"

	"github.com/JimcostDev/programming-in-go/go-exercises/exercises"
)

var filename string = "./files/txt/tabla.txt"

// CREAR UN ARCHIVO
func Save() {
	text := exercises.Iterar()
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error al crear el archivo: ", err)
		return // si hay un error, se sale de la función
	}
	fmt.Fprintln(file, text)
	file.Close()
}

// CONCATENAR TEXTO
func Concat() {
	text := exercises.Iterar()
	if !Append(text) {
		fmt.Println("Error al concatenar el texto")
	}
}

// AÑADIR TEXTO
func Append(text string) bool {
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Error al abrir el archivo: ", err)
		return false
	}
	fmt.Fprintln(file, text)
	file.Close()
	return true
}

// LEER ARCHIVO
func ReadFile() {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error al abrir el archivo: ", err)
		return
	}
	defer file.Close() // defer se ejecuta al final de la función

	data, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Error al leer el archivo: ", err)
		return
	}
	fmt.Println(string(data))
}
