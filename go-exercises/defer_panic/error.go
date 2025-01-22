package defer_panic

import (
	"errors"
	"fmt"
	"strconv"
)

// Función que realiza la división
func Division(a, b float64) (float64, error) {
	// Validar división por cero
	if b == 0 {
		return 0, errors.New("no puedes dividir entre cero")
	}
	return a / b, nil
}

// Función auxiliar para convertir cadenas a números
func ConvertirEntrada(input string) (float64, error) {
	num, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return 0, errors.New("debes ingresar un número válido")
	}
	return num, nil
}

func Error() {
	// Solicitar entradas al usuario
	var numero1Str, numero2Str string
	fmt.Print("Introduce el primer número: ")
	fmt.Scanln(&numero1Str)
	fmt.Print("Introduce el segundo número: ")
	fmt.Scanln(&numero2Str)

	// Convertir las entradas
	numero1, err := ConvertirEntrada(numero1Str)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	numero2, err := ConvertirEntrada(numero2Str)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Realizar la división
	resultado, err := Division(numero1, numero2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("El resultado de la división es: %.2f\n", resultado)
	}
}
