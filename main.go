package main

import "github.com/pedroalvaroccoyllocondori/POO_go/funciones"

func ejecutar(nombre string, funcion func(string)) {
	funcion(nombre)
}

func main() {

	nombre := "golan en español"
	//envolviendo las funciones en una variable
	saludarConLog := funciones.MiddlewareLog(funciones.Saludar)
	despedirseConLog := funciones.MiddlewareLog(funciones.Despedirse)

	//imprimiendo los valores
	ejecutar(nombre, saludarConLog)
	ejecutar(nombre, despedirseConLog)
}
