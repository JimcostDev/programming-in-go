package exercises

import (
	"fmt"
	"strconv"
)

func verificarMayorA100(s string) (bool, int, error) {
	// convertir strig a int
	num, err := strconv.Atoi(s)
	if err != nil {
		return false, 0, err //
	}

	// verificar si es mayor a 100
	esMayor := num > 100
	return esMayor, num, nil
}

func RetornarMultiplesValores() {
	resultado, numero, err := verificarMayorA100("150")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("Número: %d, Mayor a 100: %t\n", numero, resultado)
}
