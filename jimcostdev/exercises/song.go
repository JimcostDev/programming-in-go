package exercises

import (
	"bufio"
	"fmt"
	"os"
)

var songRecommendations = map[string]string{
	"Instrumental": "Mi Corazón Encantado - Dragon Ball",
	"Electronica":  "Young Again - Hardwell",
	"Clásica":      "Si Tú Me Amas - Il Divo",
	"Balada":       "Tu Carcel - Marco Antonio Solis",
	"Salsa":        "Amor de Primavera - Conjunto Chaney",
	"Bachata":      "Darte un Beso - Prince Royce",
	"Cumbia":       "Hoy - Valentino Merlo",
	"Vallenato":    "Los Caminos de la Vida - Los Diablitos",
	"Pop":          "All of Me - John Legend",
	"Romantica":    "Te Amo - Ismir Muñoz",
	"Mexicana":     "Costumbres - Natalia Jiménez",
}

func recommendSong(genre string) {
	if song, exists := songRecommendations[genre]; exists {
		fmt.Printf("Recomendación: Escucha '%s'\n", song)
	} else {
		fmt.Println("Género no encontrado")
	}
}

func ListenSong() {
	genres := []string{
		"Instrumental", "Electronica", "Clásica", "Balada", "Salsa",
		"Bachata", "Cumbia", "Vallenato", "Pop", "Romantica", "Mexicana",
	}

	fmt.Println("\nSelecciona un género musical:")
	for i, genre := range genres {
		fmt.Printf("%d. %s\n", i+1, genre)
	}

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("\nOpción: ")
	scanner.Scan()
	option := scanner.Text()

	index := map[string]string{
		"1": "Instrumental", "2": "Electronica", "3": "Clásica",
		"4": "Balada", "5": "Salsa", "6": "Bachata",
		"7": "Cumbia", "8": "Vallenato", "9": "Pop",
		"10": "Romantica", "11": "Mexicana",
	}

	if genre, exists := index[option]; exists {
		recommendSong(genre)
	} else {
		fmt.Println("Opción no válida")
	}
}
