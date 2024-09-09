// operaciones aritmeticas basicas
package exercises

import "fmt"

func Calc() {

	num1 := 5
	num2 := 8

	addition := num1 + num2
	substraction := num1 - num2
	multiplication := num1 * num2
	division := float32(num1) / float32(num2)

	fmt.Printf("La suma de %v y %v es %v\n", num1, num2, addition)
	fmt.Printf("La resta de %v y %v es %v\n", num1, num2, substraction)
	fmt.Printf("La multiplicación de %v y %v es %v\n", num1, num2, multiplication)
	fmt.Printf("La división de %v y %v es %v\n", num1, num2, division)

}
