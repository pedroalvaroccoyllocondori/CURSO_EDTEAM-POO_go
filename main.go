package main

import "fmt"

//creacion de interfaz con metodo
type Saludador interface {
	Saludador()
}

//creacion de interfase Despedidor
type Despedidor interface {
	Despedidor()
}

//embeber interfases en una sola
type SaludadorDespedidor interface {
	Saludador
	Despedidor
}

//creacion de la estructura persona
type Persona struct {
	Nombre string
}

// implemetacion de la interfaz
func (referencia Persona) Saludador() {
	fmt.Printf("hola soy una %s \n", referencia.Nombre)
}
func (referencia Persona) Despedidor() {
	fmt.Printf("adios soy un %s \n", referencia.Nombre)
}

//cracion del  tipo Texto a partir de string
type Texto string

// implementacion de la intefaz
func (referencia Texto) Saludador() {
	fmt.Printf("hola soy un %s \n ", referencia)

}
func (referencia Texto) Despedidor() {
	fmt.Printf("adios soy un %s \n", referencia)

}

//funcion que implemeta la interfaz embebida
func SaludadorDespedidorTodos(referenciass ...SaludadorDespedidor) {

	for _, valor := range referenciass {
		valor.Saludador()
		valor.Despedidor()
	}
}

func main() {
	estructura := Persona{Nombre: "alvaro"}

	var texto Texto = "Juan"

	SaludadorDespedidorTodos(estructura, texto)

}
