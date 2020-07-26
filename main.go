package main

import "fmt"

//creacion de interfaz con metodo
type Saludador interface {
	Saludador()
}

//creacion de una estructura
type Persona struct {
	Nombre string
}

// implemetacion de la interfaz
func (referencia Persona) Saludador() {
	fmt.Printf("hola soy una %s", referencia.Nombre)
}

//cracion de otro tipo
type Texto string

// implementacion de la intefaz
func (referencia Texto) Saludador() {
	fmt.Printf("hola soy un %s ", referencia)

}

func main() {
	var estructura Saludador = Persona{Nombre: "alvaro"}
	estructura.Saludador()

	fmt.Println("")

	var texto Saludador = Texto("Juan")
	texto.Saludador()

}
