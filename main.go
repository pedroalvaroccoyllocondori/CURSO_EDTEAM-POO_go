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

func SaludadoraTodos(referenciass ...Saludador) { //esta funcion solo permite almacenar los
	//tipos de datos que tengan esta interfaz
	for _, valor := range referenciass {
		valor.Saludador()
		fmt.Printf("\t valor :%v , tipo :%T \n", valor, valor)
	}
}

func main() {
	estructura := Persona{Nombre: "alvaro"}

	var texto Texto = "Juan"

	SaludadoraTodos(estructura, texto)

}
