package funciones

import (
	"fmt"
	"time"
)

// funcion decoradora que  recibe una funcion le agrega  funcionalidad  y devuelve una funcion
func MiddlewareLog(funcion func(string)) func(string) {
	return func(nombre string) {
		fmt.Println(time.Now().Format("2006-01-02 15:04:05")) // funcionalidad agregada a la funcion
		funcion(nombre)
	}
}
