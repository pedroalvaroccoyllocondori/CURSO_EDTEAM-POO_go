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

func (instancia curso) imprimirclases(){// metodo perteneciente a la clase curso
	texto:="las clases son : "
	for _, valor := range instancia.clases{
		texto+= valor +", "
	}
	fmt.Println(texto[:len(texto)-2])//metodo para borrar la ultima coma en el programa
}

func main(){
	//instancia especificando los atributos
	 Go := curso{
		 nombre:"javascrip desde cero",
		 precio:68.05,
		 esgratis:false,
		 idusuarios:[]uint{01,02,03},
		 clases: map[uint]string{
			 1:"introducion",
			 2:"estructuras",
			 3:"mapas",
		 },	 
	 }
	 
	 Go.imprimirclases()

}