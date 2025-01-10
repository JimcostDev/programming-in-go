package main

import "fmt"

func main() {
	ciudades := []string{"Tokyo", "Lleida", "Paris", "Madrid"}

	fmt.Println("Primera forma:")
	for i := 0; i < len(ciudades); i++ {
		fmt.Printf("[%d] %s\n", i, ciudades[i])
	}

	fmt.Println("Segunda forma:")
	for i, ciudad := range ciudades {
		fmt.Printf("[%d] %s\n", i, ciudad)
	}

	fmt.Println("Tercera forma:")
	for _, ciudad := range ciudades {
		fmt.Println(ciudad)
	}
}
