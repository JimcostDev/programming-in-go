package arrays_slices

import "fmt"

func Equipos() {
	// Crear un slice de equipos
	equipos := []string{"Deportivo Cali 💚", "Liverpool ❤️", "Real Madrid 🤍"}

	// Agregar más equipos al slice
	equipos = append(equipos, "Boca Juniors 💙💛", "Arsenal 💖")

	// Crear un slice basado en otro (subslice)
	favoritos := equipos[0:3] // Incluye los índices 1 a 3 (excluye el índice 4)

	// Modificar un elemento del slice original
	equipos[4] = "Barcelona ❤️💙"

	// Imprimir los slices
	fmt.Println("Equipos:", equipos)
	fmt.Println("Favoritos:", favoritos)

	// Ver longitud y capacidad del slice
	fmt.Printf("Longitud de 'equipos': %d, Capacidad: %d\n", len(equipos), cap(equipos))

	// ver tipo de 'equipos'
	fmt.Printf("Tipo: %T\n", equipos) // al tener un tamaño dinámico, lo identifica como un slice
}
