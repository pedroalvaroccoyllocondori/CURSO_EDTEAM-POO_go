package main

import (
	"fmt"
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
}
