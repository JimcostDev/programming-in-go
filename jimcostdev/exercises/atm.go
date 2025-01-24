/*
¿Cómo entrega dinero un cajero automático?

Puede que hayas ido varias veces a un cajero, y a pesar de que le pides la misma
cantidad de dinero puede ser que te entregue el monto en diferentes
combinaciones de billetes, a veces te entrega todos del mismo valor,
a veces diferentes. ¿Cómo sabe el cajero cuál es la combinación de billetes adecuada?
Crearemos un algoritmo que responda esa pregunta.
*/
package exercises

import (
	"fmt"
	"strings"
)

// total disponible en el cajero = 1350
var disponible = map[int]int{
	100: 3,
	50:  6,
	20:  10,
	10:  40,
	5:   10,
	2:   30,
	1:   40,
}

// es un slice de enteros que contiene los valores de los billetes/monedas que el cajero puede entregar
var denominaciones = []int{100, 50, 20, 10, 5, 2, 1}

func ATM(dineroARetirar int) {
	fmt.Println("Cantidad disponible: ", calcularTotalDisponible())
	valorARetirar := dineroARetirar // valor a retirar
	fmt.Println("Valor a retirar: ", valorARetirar)
	// - * 10 + "\n" es una forma de imprimir una línea de guiones
	fmt.Println(strings.Repeat("-", 10))

	if valorARetirar > calcularTotalDisponible() {
		fmt.Println("No hay suficiente dinero en el cajero, intente con una cantidad menor")
		return
	}

	saldoARetirar := valorARetirar
	for saldoARetirar > 0 {
		for _, denominacion := range denominaciones {
			cantidad := obtenerCantidadPorDenominacion(denominacion, saldoARetirar)
			saldoARetirar -= cantidad * denominacion
			if cantidad > 0 {
				fmt.Printf("Entregar %d billetes de %d\n", cantidad, denominacion)
			}
		}
	}
}

// función que calcula el saldo disponible
func calcularTotalDisponible() int {
	total := 0
	for billete, cantidad := range disponible {
		total += billete * cantidad
	}

	return total
}

// función que devuelve la cantidad de billetes o monedas por su denominación
func obtenerCantidadPorDenominacion(denominacion, saldoARetirar int) int {
	cantidad := saldoARetirar / denominacion
	if cantidad > disponible[denominacion] {
		cantidad = disponible[denominacion]
	}
	disponible[denominacion] -= cantidad

	return cantidad
}
