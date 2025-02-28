package exercises

import (
	"fmt"

	// paquete de utilidades de jimcostdev (proyecto)
	"github.com/JimcostDev/jimcostdev-utils/strutils"
)

// TitleCaseExample demuestra el uso de ToTitleCase
func TitleCaseExample(input string) {
	result := strutils.ToTitleCase(input)
	fmt.Println("Resultado en Title Case:", result)
}
