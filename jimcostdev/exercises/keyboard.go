package exercises

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var number int
var number2 int
var label string
var err error

func Keyboard() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Ingrese el primer número: ")
	if scanner.Scan() {
		number, err = strconv.Atoi(scanner.Text())
		if err != nil {
			panic("Error al convertir el número 1" + err.Error())
		}
	}

	fmt.Println("Ingrese el segundo número: ")
	if scanner.Scan() {
		number2, err = strconv.Atoi(scanner.Text())
		if err != nil {
			panic("Error al convertir el número 2" + err.Error())
		}
	}

	fmt.Println("Ingrese la etiqueta: ")
	if scanner.Scan() {
		label = scanner.Text()
	}

	fmt.Println(label, number*number2)
}
