package unittests

import "testing"

// go test . -> correr todas las pruebas
// go test -v -> correr todas las pruebas con información detallada
// go test -run TestName -> correr una prueba en específico
// go test -cover -> para ver el porcentaje de cobertura de las pruebas
// go test -timeout 3s -> para establecer un tiempo de espera para las pruebas

// TestFactorial pruebas para la función Factorial
func TestFactorial(t *testing.T) {
	f := Factorial(5)
	if f != 120 {
		t.Errorf("Factorial(5): Se esperaba 120. Retorno %d;", f)
	}
}

// TestSumarEnteros pruebas para la función SumarEnteros
func TestSumarEnteros(t *testing.T) {
	s := SumarEnteros([]int{1, 2, 3, 4, 5})
	if s != 15 {
		t.Errorf("SumarEnteros([1, 2, 3, 4, 5]): Se esperaba 15. Retorno %d;", s)
	}
}

// TestSumarEnteros pruebas para la función SumarEnteros usando table-driven tests (el enfoque estándar en Go):
func TestSumarEnterosPro(t *testing.T) {
	tests := []struct {
		nombre   string
		entrada  []int
		esperado int
	}{
		{
			nombre:   "Lista normal",
			entrada:  []int{1, 2, 3, 4},
			esperado: 10,
		},
		{
			nombre:   "Lista vacía",
			entrada:  []int{},
			esperado: 0,
		},
		{
			nombre:   "Un solo elemento",
			entrada:  []int{5},
			esperado: 5,
		},
		{
			nombre:   "Números negativos",
			entrada:  []int{-1, 2, -3},
			esperado: -2,
		},
		{
			nombre:   "Ceros",
			entrada:  []int{0, 0, 0},
			esperado: 0,
		},
	}

	for _, test := range tests {
		t.Run(test.nombre, func(t *testing.T) {
			resultado := SumarEnteros(test.entrada)
			if resultado != test.esperado {
				t.Errorf("SumarEnteros(%v) = %d, se esperaba %d", test.entrada, resultado, test.esperado)
			}
		})
	}
}
