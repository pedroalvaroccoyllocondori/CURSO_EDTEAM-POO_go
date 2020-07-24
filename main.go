package main

import (
	"fmt"
)

type curso struct {
	nombre string
}

func (instancia curso) Imprimir() {
	fmt.Printf("%+v\n", instancia)
}

//declaracion de alias de tipo
type MyAlias = curso //hereda los atributos y metodos

//definicion de tipo
type NuevoCurso curso // no hereda loa atributos y metodos

func main() {

	//declaracion de alias
	objeto := MyAlias{nombre: "java"}
	objeto.Imprimir()
	fmt.Printf("el tipo es %T\n", objeto) //tipo de dato

	// declaracion de  definicion de tipo
	objeto1 := NuevoCurso{nombre: "git"}
	//objeto1.Imprimir()
	fmt.Printf("el tipo es %T\n", objeto1) //tipo de dato

}
