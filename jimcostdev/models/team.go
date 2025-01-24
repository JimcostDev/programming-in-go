package models

import (
	"fmt"
	"time"
)

type Team struct {
	ID        int
	Name      string
	CreatedAt time.Time
}

// Constructor
func NewTeam(id int, name string) *Team {
	return &Team{
		ID:        id,
		Name:      name,
		CreatedAt: time.Now(),
	}
}

// Método para mostrar información del equipo
func (t *Team) DisplayInfo() {
	// formatear la fecha
	formattedDate := t.CreatedAt.Format("2006-01-02")
	fmt.Printf("Team ID: %d, Name: %s, Created At: %s\n", t.ID, t.Name, formattedDate)
}

// Método para actualizar el nombre
func (t *Team) UpdateName(newName string) {
	t.Name = newName
}
