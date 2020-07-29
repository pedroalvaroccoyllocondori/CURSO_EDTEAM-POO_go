package main

import (
	"fmt"
)

type ejemplo interface {
	metodo1()
}

func envoltura(interfaz interface{}) { //metodo que recibe una interfaz vacia
	fmt.Printf("valor :%v tipo:%T \n", interfaz, interfaz)
}

func main() {

	//var ejemplar ejemplo // declarando una interfaz nula (ninguna clase a implementado ese metodo)
	//ejemplar.metodo1()
	envoltura(12)
	envoltura(12.23)
	envoltura("cadena")
	envoltura(true)

}
