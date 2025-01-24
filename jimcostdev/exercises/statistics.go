package exercises

import (
	"fmt"
	"sort"
)

func Statistics() {
	lista := []int{0, 10, 0, 0, 0, 1, 1, 1, 1, 1, 1, 2, 2, 2, 2, 2, 2, 2, 2, 3, 3, 3, 3, 4, 4}
	cantidad := len(lista)

	// Crear una copia de la lista para mantener el orden original
	listaOrdenada := make([]int, cantidad)
	copy(listaOrdenada, lista)

	// ordenar la lista de elementos
	sort.Ints(listaOrdenada)

	// mostrar la lista ordenada
	fmt.Println(listaOrdenada)

	// Mostrar la cantidad total de datos
	fmt.Printf("Total de datos(N): %d\n", cantidad)

	// obtener la suma de los elementos de la lista (total)
	sumarElemetos := 0
	for _, item := range listaOrdenada {
		sumarElemetos += item
	}

	fmt.Printf("La suma de los datos es: %d\n", sumarElemetos)

	// hallar promedio
	promedio := float64(sumarElemetos) / float64(cantidad)
	fmt.Printf("El promedio es: %v\n", promedio)
}
