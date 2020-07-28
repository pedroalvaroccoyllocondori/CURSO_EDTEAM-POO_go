package main

import (
	"fmt"
)

//creacion de interfases
type SetterGetter interface {
	Obtener() string
	Establecer(string)
}

// creacion de una estructura
type Persona struct {
	nombre string
}

func NuevaPersona(nombre string) Persona {
	return Persona{nombre}

}
func (referencia Persona) Obtener() string {
	return referencia.nombre
}
func (referencia *Persona) Establecer(nombre string) { //receptor de tipo puntero para actualizar la informacion
	referencia.nombre = nombre
}

// funcion que  que recibe una instancia que tiene los metodos de la intefaz
func Executar(instancia SetterGetter, nombre string) {
	instancia.Establecer(nombre)
	fmt.Println(instancia.Obtener())
}
func main() {

	persona := NuevaPersona("juan") // instanciando un objeto

	//Executar(persona, "alvaro") // no se puede ejecutara ya qq el emtodo tipo puntero no implemta todos losa metodos de la intefaz
	Executar(&persona, "alvaro")
}
