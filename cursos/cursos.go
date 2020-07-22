package main

import(
	"fmt"
)

type curso struct{
	nombre string
	precio float64
	esgratis bool
	idusuarios []uint
	clases map[uint]string


}

func (instancia *curso) imprimirclases(){// metodo perteneciente a la clase curso
	texto:="las clases son : "
	for _, valor := range instancia.clases{
		texto+= valor +", "
	}
	fmt.Println(texto[:len(texto)-2])//metodo para borrar la ultima coma en el programa
}

func (instancia *curso) cambiarprecio(precio float64){
	// cuando se requieren actualizar datos en la  instancia
	// se utiliza  llos punteros (*) delante de  la clase de referncia (*curso)
	instancia.precio=precio
}