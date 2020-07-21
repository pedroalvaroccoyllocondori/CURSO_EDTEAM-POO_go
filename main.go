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
	 //instancia sin especificar atributos
	 Css := curso{
		// no es necesario incluir ellnombre de los campos
		"css desde cero",
		68.05,
		false,
		[]uint{01,02,03},
	    map[uint]string{
			1:"introducion",
			2:"estructuras",
			3:"mapas",
		},	 
	}
// instancia de una linea
	java :=curso{nombre:"java",esgratis:true}
// instancia sin argumentos
	python :=curso{}
	python.nombre="phthon desde cero"
	python.idusuarios=[]uint{01,02}


	fmt.Println(Go.nombre)
	fmt.Println(Css.nombre)
	fmt.Printf("%+v\n",java)
	fmt.Printf("%+v\n",python)

	

}