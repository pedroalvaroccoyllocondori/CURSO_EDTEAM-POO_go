package main

import (
	"fmt"
	"strings"
)

type ejemplo interface {
	metodo1()
}

func envoltura(interfaz interface{}) { //metodo que recibe una interfaz vacia

	switch valor := interfaz.(type) {
	case string:
		fmt.Printf("\t%s\n", strings.ToUpper(valor))
	case int:
		fmt.Println(valor * 5)
	default:
		fmt.Printf("valor :%v tipo:%T \n", interfaz, interfaz)
		//validar el tipo de  interfaz que ingresa al metodo

	}
}
func main() {

	//var ejemplar ejemplo // declarando una interfaz nula (ninguna clase a implementado ese metodo)
	//ejemplar.metodo1()
	envoltura(12)
	envoltura("matusculas")
	envoltura(12.23)
	envoltura("cadena")
	envoltura(true)

}
