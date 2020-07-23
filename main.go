package main

import (
	"fmt"

	"github.com/pedroalvaroccoyllocondori/POO_go"
)

func main() {
	//instancia especificando los atributos
	Go := cursos.NuevoCurso("go desde cero", 0, false)

	Go.Estableceridusuario([]uint{10, 100, 10})
	Go.Establecerclases(map[uint]string{
		1: "introduccion",
		2: "estructuras",
		3: "mapas",
	})
	Go.Establecernombre("go programacion orientada a objetos")
	fmt.Println(Go.Obtenernombre())

	fmt.Printf("%+v", Go)

	//utilizar los geter  y setter

}
