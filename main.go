package main

import "github.com/pedroalvaroccoyllocondori/POO_go/funciones"

func ejecutar(nombre string, funcion func(string)) {
	funcion(nombre)
}

func main() {

	nombre := "golan en español"
	ejecutar(nombre, funciones.Saludar)
	ejecutar(nombre, funciones.Despedirse)
}
