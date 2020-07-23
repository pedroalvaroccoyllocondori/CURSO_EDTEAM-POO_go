package main

import (
	"fmt"

	"github.com/pedroalvaroccoyllocondori/POO_go/pkg/cliente"
	"github.com/pedroalvaroccoyllocondori/POO_go/pkg/factura"
	"github.com/pedroalvaroccoyllocondori/POO_go/pkg/idfactura"
)

func main() {

	objeto := factura.Nuevo(
		"peru",
		"ica",
		cliente.Nuevo("alvaro", "ica", "917186613"),
		[]idfactura.Item{
			idfactura.Nuevo(01, "curso de git", 12.3),
			idfactura.Nuevo(02, "curso de android", 12.3),
			idfactura.Nuevo(03, "curso de go", 67.3),
		},
	)
	objeto.Establecertotal()
	fmt.Printf("%+v", objeto)
}
