package main

import "fmt"

// creacion de la clase  humano

type Humano struct {
	Edad  uint
	Hijos uint
}

func NuevoHumano(edad, hijos uint) Humano {
	return Humano{edad, hijos}
}

// creacion de la  clase persona
type Persona struct {
	Nombre string
	Edad   uint
}

func NuevaPersona(nombre string, edad uint) Persona {
	return Persona{nombre, edad}
}

func (referncia Persona) Saludo() {
	fmt.Println("hola a todos desde persona")
}

// creacion de la estructura empleado a partir de persona
type Empleado struct {
	Persona // embebemos  la clase persona y automaticamente  hereda todo los atributos y metodos
	// el receptor de esos metodos sera el tipo interno
	Humano // embebos la clase empleado parecido a  la herencia  multiple
	sueldo float64
}

func NuevoEmpleado(nombre string, edad uint, hijos uint, sueldo float64) Empleado {
	return Empleado{NuevaPersona(nombre, edad), NuevoHumano(edad, hijos), sueldo}
}
func (instancia Empleado) Honorario() {
	fmt.Println(instancia.sueldo * 0.9)
}

//sobre escritura de metodo (en go no existe sobreescritura solo ocultamiento de metodos)
func (instancia Empleado) Saludo() { //
	fmt.Println("hola desde empleado")

}

func main() {

	empleado := NuevoEmpleado("juan", 25, 5, 960)
	fmt.Println(empleado.Nombre, empleado.Humano.Edad) // referencia expplicita  para ver a que campos quiere q acceda  el programa
	empleado.Saludo()                                  // verifica si empleado tiene ese metodo  si lo tiene  lo ejecuta
	// si no lo tiene  va al siguiente nivel de embebimiento
	empleado.Persona.Saludo() // si quermos saltar el ocultamiento de metodos  solo tenemos que accedere
	// a ese metodo  desde la estructura persona
	empleado.Honorario()

}
