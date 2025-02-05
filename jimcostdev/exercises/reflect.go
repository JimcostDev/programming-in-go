package exercises

// Reflection es la habilidad de un programa de inspeccionar los tipos y valores de las variables en tiempo de ejecución
import (
	"fmt"
	"reflect"
)

func Reflection() {
	me := jimcostdev{name: "Ronaldo", color: "Verde", team: "Cali"}
	meType := reflect.TypeOf(me)   //  tipo exacto de la variable, en este caso jimcostdev
	meKind := meType.Kind()        // "tipo base" o categoría del tipo, en este caso struct
	meValue := reflect.ValueOf(me) // valor almacenado en la variable

	fmt.Println(meType)
	fmt.Println(meKind)
	fmt.Println(meValue)

	age := 25
	fmt.Println(reflect.TypeOf(age))
	fmt.Println(reflect.TypeOf(age).Kind())
	fmt.Println(reflect.ValueOf(age))
}

type jimcostdev struct {
	name  string
	color string
	team  string
}
