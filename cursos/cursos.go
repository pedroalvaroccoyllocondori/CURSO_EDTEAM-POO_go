package cursos

import "fmt"

type curso struct {
	// los atributos que empiezan  la primera leta en mayuscula
	//significa qu es exportado

	nombre     string
	precio     float64
	esgratis   bool
	idusuarios []uint
	clases     map[uint]string // esta variable sera privada (ya que no tiene los modificadores de acceso)
}

// funcion NuevoCurso en luagar de metodo constructor

func NuevoCurso(nombre string, precio float64, esgratis bool) *curso {
	if precio == 0 {
		precio = 50

	}
	return &curso{
		nombre:   nombre,
		precio:   precio,
		esgratis: esgratis,
	}

}

//metodos de setter
func (instancia *curso) Establecernombre(nombre string) {
	instancia.nombre = nombre
}
func (instancia *curso) Establecerprecio(precio float64) {
	instancia.precio = precio
}
func (instancia *curso) Establecergratis(esgratis bool) {
	instancia.esgratis = esgratis
}
func (instancia *curso) Estableceridusuario(idusuarios []uint) {
	instancia.idusuarios = idusuarios
}
func (instancia *curso) Establecerclases(clases map[uint]string) {
	instancia.clases = clases
}

//metodos getter
func (instancia *curso) Obtenernombre() string {
	return instancia.nombre
}
func (instancia *curso) Obtenerprecio() float64 {
	return instancia.precio
}
func (instancia *curso) Obteneresgratis() bool {
	return instancia.esgratis
}
func (instancia *curso) Obteneridusuario() []uint {
	return instancia.idusuarios
}
func (instancia *curso) Obtenerclases() map[uint]string {
	return instancia.clases
}

//funcion modificada debido  a que  no tiene el modifiador  de acceso
func (instancia *curso) Imprimirclases() { // metodo perteneciente a la clase curso
	texto := "las clases son : "
	for _, valor := range instancia.clases {
		texto += valor + ", "
	}
	fmt.Println(texto[:len(texto)-2]) //metodo para borrar la ultima coma en el programa
}
