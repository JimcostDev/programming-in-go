package arrays_slices

import "fmt"

func Equipos() {
	// Crear un slice de equipos
	equipos := []string{"Deportivo Cali 💚", "Liverpool ❤️", "Real Madrid 🤍"}

	// Agregar más equipos al slice
	equipos = append(equipos, "Boca Juniors 💙💛", "Arsenal 💖")
	equipos = append(equipos, "River Plate 🤍❤️", "Chelsea 💙")

	// Crear un slice basado en otro (subslice)
	favoritos := equipos[0:3] // Incluye los índices 0, 1 y 2, excluye el 3
	// Modificar un elemento del slice original
	equipos[4] = "Barcelona ❤️💙"

	// Imprimir los slices
	for _, equipo := range equipos {
		fmt.Println(equipo)
	}
	fmt.Println("Favoritos:", favoritos)

	// Ver longitud y capacidad del slice
	fmt.Printf("Longitud de 'equipos': %d, Capacidad: %d\n", len(equipos), cap(equipos))
	// capacidad inicial es de 3, pero al agregar más elementos, se duplica a 6, luego a 12, si se sigue agregando, se duplica nuevamente

	// ver tipo de 'equipos'
	fmt.Printf("Tipo: %T\n", equipos) // al tener un tamaño dinámico, lo identifica como un slice

	// uso de ... para desempaquetar
	// Definimos dos slices
	numeros1 := []int{1, 2, 3}
	numeros2 := []int{4, 5, 6}

	// Usamos `append` con `...` para combinar ambos slices
	resultado := append(numeros1, numeros2...)
	// ... expande los elementos del slice resultante como
	// argumentos individuales para la función append, Esto es necesario porque append
	// espera una lista de elementos, no otro slice.

	// Mostramos el resultado
	fmt.Println(resultado) // Salida: [1 2 3 4 5 6]
}
