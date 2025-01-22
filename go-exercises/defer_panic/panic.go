package defer_panic

import "fmt"

func dividir(a, b int) int {
	if b == 0 {
		panic("No se puede dividir entre cero")
	}
	return a / b
}

func Panic() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Error recuperado:", r) // se muestra el error al final
		}
	}()

	fmt.Println("Resultado:", dividir(10, 2)) // 10 / 2 = 5
	fmt.Println("Resultado:", dividir(10, 0)) // Panic aquí
	fmt.Println("Esta línea no se ejecuta")
}
