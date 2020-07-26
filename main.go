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

//funcion que solo funciona con la interfase saludador()
func SaludadoraTodos(referenciass ...Saludador) { //esta funcion solo permite almacenar los
	//tipos de datos que tengan esta interfaz
	for _, valor := range referenciass {
		valor.Saludador()
	}
}
func DespedidoraaTodos(referenciass ...Despedidor) { //esta funcion solo permite almacenar los
	//tipos de datos que tengan esta interfaz
	for _, valor := range referenciass {
		valor.Despedidor()
	}
}

func main() {
	estructura := Persona{Nombre: "alvaro"}

	var texto Texto = "Juan"

	SaludadoraTodos(estructura, texto)
	DespedidoraaTodos(estructura, texto)

}
