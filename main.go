package main

import "github.com/pedroalvaroccoyllocondori/POO_go/funciones"

type myFuncion func(string) //creando un tipo de dato de tipo mi funcion

func ejecutar(nombre string, funcion funciones.Mifuncion) {
	funcion(nombre)
}

func main() {

	nombre := "golan en español"
	//envolviendo las funciones en una variable
	//imprimiendo los valores
	ejecutar(nombre, funciones.MiddlewareLog(funciones.Saludar))
	ejecutar(nombre, funciones.MiddlewareLog(funciones.Despedirse))
}
