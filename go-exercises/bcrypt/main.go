package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	// Mensaje a encriptar
	password := "Cambiar123*!"

	// Encriptar el mensaje
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	fmt.Println(hashedPassword)

	// Cambiar el mensaje - contraseña incorrecta
	// password = "Cambiar123*"
	// Comparar un mensaje con el hash almacenado
	err = bcrypt.CompareHashAndPassword(hashedPassword, []byte(password))
	if err != nil {
		fmt.Println("Contraseña incorrecta")
	} else {
		fmt.Println("Contraseña correcta")
	}
}
