package exercises

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func Iterar() {
	// tabla de multiplicar del 1 al 10 de un número ingresado por el usuario
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("Ingresa un número entero: ")
		if scanner.Scan() {
			number, err = strconv.Atoi(scanner.Text()) // utilizamos las variables declaradas en keyboard.go
			if err != nil {
				fmt.Println("Error al convertir el número: ", err)
				continue // si hay un error, se vuelve a pedir el número
			} else {
				break // si no hay error, se sale del ciclo
			}
		}
	}

	// ciclo for para imprimir la tabla de multiplicar
	for i := 1; i <= 10; i++ {
		fmt.Println(number, "x", i, "=", number*i)
	}

}
