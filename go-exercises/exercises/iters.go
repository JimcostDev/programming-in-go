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
				continue
			}
			// Si llegamos aquí, el número es válido, así que imprimimos la tabla
			for i := 1; i <= 10; i++ {
				fmt.Printf("%d x %d = %d\n", number, i, number*i)
			}
			break // Salimos del bucle principal
		} else {
			fmt.Println("Error al leer el número: ", scanner.Err())
			break
		}
	}

}
