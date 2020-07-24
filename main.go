package main

import (
	"fmt"
<<<<<<< HEAD
)

type NuevoBoleano bool //definiciones de nuevos tipos  con base en tipos predeclarados

func (referencia NuevoBoleano) Cadena() string {
	if referencia {
		return "VERDADERO"
	} else {
		return "FALSO"
	}
}

func main() {

	var boleano NuevoBoleano = false

	fmt.Println(boleano.Cadena())
=======

	"github.com/pedroalvaroccoyllocondori/POO_go/pkg/cliente"
	"github.com/pedroalvaroccoyllocondori/POO_go/pkg/factura"
	"github.com/pedroalvaroccoyllocondori/POO_go/pkg/idfactura"
)

func main() {

	objeto := factura.Nuevo(
		"peru",
		"ica",
		cliente.Nuevo("alvaro", "ica", "917186613"),

		idfactura.NuevoItems(
			idfactura.Nuevo(01, "curso de git", 12.3),
			idfactura.Nuevo(02, "curso de android", 12.3),
			idfactura.Nuevo(03, "curso de go", 67.3),
		),
	)
	objeto.Establecertotal()
	fmt.Printf("%+v", objeto)
>>>>>>> desarrollo
}
