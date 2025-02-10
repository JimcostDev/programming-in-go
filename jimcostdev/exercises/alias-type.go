package exercises

import "fmt"

//Tipos personalizados en Go
type Kage string

const (
	Hokage     Kage = "Hokage"
	Mizukage   Kage = "Mizukage"
	Raikage    Kage = "Raikage"
	Kazekage   Kage = "Kazekage"
	Tsuchikage Kage = "Tsuchikage"
)

func Types() {
	var kakashi Kage = Hokage
	fmt.Println("Kakashi es un:", kakashi)
}
