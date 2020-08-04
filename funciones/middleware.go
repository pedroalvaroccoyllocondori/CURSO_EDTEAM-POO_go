package funciones

import (
	"fmt"
	"time"
)

//tipo de dato funcion
type Mifuncion func(string)

// funcion decoradora que  recibe una funcion le agrega  funcionalidad  y devuelve una funcion

// funcion que devuelve un tipo de dato Mifuncion
func MiddlewareLog(funcion func(string)) Mifuncion {
	return func(nombre string) {
		fmt.Println("inicio", time.Now().Format("2006-01-02 15:04:05")) // funcionalidad agregada a la funcion
		funcion(nombre)
		fmt.Println("fin", time.Now().Format("2006-01-02 15:04:05")) // funcionalidad agregada a la funcion

	}
}
