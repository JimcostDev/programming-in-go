package arrays_slices

import "fmt"

func Teams() {
	// var teams [3]string
	// teams[0] = "Deportivo Cali 💚"
	// teams[1] = "Liverpool ❤️"
	// teams[2] = "Real Madrid 💙"

	// [...] indica que el tamaño del array es fijo, pero se deduce automáticamente a partir de los elementos proporcionados en la inicialización.
	teams := [...]string{"Deportivo Cali 💚", "Liverpool ❤️", "Real Madrid 🤍"}

	// fmt.Println(teams)
	for _, team := range teams {
		fmt.Println(team)
	}

	// Imprimir el tipo de 'teams'
	fmt.Printf("Tipo: %T\n", teams) // al tener un tamaño fijo, lo identifica como un array
}
