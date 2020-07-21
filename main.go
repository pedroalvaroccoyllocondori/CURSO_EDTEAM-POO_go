package main







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