/*
Dada una lista de por lo menos 3 participantes se deberá retornar
un cronograma de enfrentamientos de tal manera que todos los
equipos se enfrenten contra cada uno de ellos de manera aleatoria
y sin repetición para poder tener el calendario de partidos de un campeonato.
*/
package exercises

import (
	"fmt"
)

var participantes = []string{"Cali", "America", "Nacional", "Millonarios", "Medelllin"}

func Tournament() {
	for _, v := range sorteo(participantes) {
		fmt.Println(v)
	}
}

func sorteo(equipos []string) [][]string {
	if len(equipos)%2 != 0 {
		equipos = append(equipos, "Descansa")
	}

	var fixture [][]string
	partidos := len(equipos) / 2
	rivales := len(equipos) - 1

	for i := 0; i < rivales; i++ {
		fixture = append(fixture, []string{})

		for j := 0; j < partidos; j++ {
			fixture[i] = append(fixture[i], []string{equipos[j], equipos[rivales-j]}...)
		}

		lastIndex := len(equipos) - 1
		var temp []string
		temp = append(temp, equipos[0], equipos[lastIndex])
		temp = append(temp, equipos[1:lastIndex]...)

		equipos = temp
	}

	return fixture
}
