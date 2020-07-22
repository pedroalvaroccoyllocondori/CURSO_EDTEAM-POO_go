package main

import (
	"fmt"

	"github.com/pedroalvaroccoyllocondori/POO_go/cursos/cursos"
)

func main() {
	//instancia especificando los atributos
	Go := &cursos.Curso{
		Nombre:     "javascrip desde cero",
		Precio:     68.05,
		Esgratis:   false,
		Idusuarios: []uint{01, 02, 03},
		Clases: map[uint]string{
			1: "introducion",
			2: "estructuras",
			3: "mapas",
		},
	}
	(Go).Cambiarprecio(100.66)
	fmt.Println(Go.Precio)

}
