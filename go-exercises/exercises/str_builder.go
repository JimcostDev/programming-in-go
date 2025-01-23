package exercises

import (
	"fmt"
	"strings"
)

func StrBuilder() {
	// strings.Builder es una estructura que permite concatenar strings de manera eficiente
	// en Go, ya que no se crean nuevas cadenas de texto en cada concatenación
	var sb strings.Builder
	sb.WriteString("Hola ")
	sb.WriteString("Mundo")
	fmt.Println(sb.String())
}
