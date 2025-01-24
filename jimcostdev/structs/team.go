package structs

import (
	"github.com/JimcostDev/programming-in-go/jimcostdev/models"
)

func Team() {
	// Crear un nuevo equipo
	team := models.NewTeam(1, "Deportivo Cali")

	// Mostrar información
	team.DisplayInfo()

	// Actualizar el nombre
	team.UpdateName("Liverpool")

	// Mostrar información actualizada
	team.DisplayInfo()
}
