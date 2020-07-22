package cursos

import "fmt"

type curso struct {
	// los atributos que empiezan  la primera leta en mayuscula
	//significa qu es exportado

	Nombre     string
	Precio     float64
	Esgratis   bool
	Idusuarios []uint
	Clases     map[uint]string
}

// funcion new en luagar de metodo constructor

func NuevoCurso(nombre string, precio float64, esgratis bool) *curso {
	if precio == 0 {
		precio = 50

	}
	return &curso{
		Nombre:   nombre,
		Precio:   precio,
		Esgratis: esgratis,
	}

}

func (instancia *curso) Imprimirclases() { // metodo perteneciente a la clase curso
	texto := "las clases son : "
	for _, valor := range instancia.Clases {
		texto += valor + ", "
	}
	fmt.Println(texto[:len(texto)-2]) //metodo para borrar la ultima coma en el programa
}

func (instancia *curso) Cambiarprecio(Precio float64) {
	// cuando se requieren actualizar datos en la  instancia
	// se utiliza  llos punteros (*) delante de  la clase de referncia (*curso)
	instancia.Precio = Precio
}
