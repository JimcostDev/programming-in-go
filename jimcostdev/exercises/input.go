// ejemplo de entrada de datos - Scan y Scanf
package exercises

import "fmt"

func Input() {

	// Scanf - permite especificar el formato de entrada con mas detalle, en este caso se espera que el usuario ingrese la hora en formato HH:MM:SS
	var hour, minute, second int
	fmt.Println("HH:MM:SS? ")
	fmt.Scanf("%d:%d:%d", &hour, &minute, &second)

	// Scan - permite leer un valor de entrada
	var user string
	fmt.Println("Ingresa tu usuario:")
	fmt.Scan(&user)

	fmt.Printf("Hola %v, bienvenido!\n", user)
	fmt.Printf("La hora ingresada es: %v:%v:%v\n", hour, minute, second)
}
