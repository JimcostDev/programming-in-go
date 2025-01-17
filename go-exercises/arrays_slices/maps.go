package arrays_slices

import (
	"fmt"
)

func Maps() {
	// Crear un mapa que asocia nombres de equipos con sus colores
	equipos := map[string]string{
		"Deportivo Cali": "Verde 💚",
		"Liverpool":      "Rojo ❤️",
		"Real Madrid":    "Blanco 🤍",
	}

	// Acceder a un valor en el mapa
	fmt.Println("Color de Deportivo Cali:", equipos["Deportivo Cali"])

	// Agregar un nuevo equipo al mapa
	equipos["Barcelona"] = "Grana y Azul 💙❤️"

	// Verificar si una clave existe en el mapa
	color, existe := equipos["Boca Juniors"]
	if existe {
		fmt.Println("Color de Boca Juniors:", color)
	} else {
		fmt.Println("Boca Juniors no está en el mapa.")
	}

	// Recorrer el mapa
	fmt.Println("Equipos y sus colores:")
	for equipo, color := range equipos {
		fmt.Printf("- %s: %s\n", equipo, color)
	}

	// Eliminar un equipo del mapa
	delete(equipos, "Real Madrid")

	// Imprimir el mapa actualizado
	fmt.Println("Mapa actualizado:", equipos)
}
