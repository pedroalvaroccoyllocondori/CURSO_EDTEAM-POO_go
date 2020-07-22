package main

import (
	"fmt"

	"github.com/pedroalvaroccoyllocondori/POO_go/cursos/cursos"
)

func main() {
	//instancia especificando los atributos
	Go := cursos.NuevoCurso("go desde cero", 0, false)
	Go.Idusuarios = []uint{10, 100, 10}
	Go.Clases = map[uint]string{
		1: "introduccion",
		2: "estructuras",
		3: "mapas",
	}

	Go.Imprimirclases()
	fmt.Printf("%+v", Go)

}
