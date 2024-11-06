package exercises

import (
	"fmt"
	"runtime"
)

func Run() {
	// Obtener el sistema operativo en ejecución
	os := runtime.GOOS

	if os == "linux" || os == "darwin" {
		fmt.Println("SO: ", os)
	} else {
		fmt.Println("SO: ", os)
	}

	fmt.Print("Estás ejecutando Go en: ")
	// ejemplo de switch
	switch os {
	case "windows":
		fmt.Println("Windows")
	case "darwin":
		fmt.Println("macOS")
	case "linux":
		fmt.Println("Linux")
	default:
		fmt.Printf("%s (no identificado)", os)
	}
}
