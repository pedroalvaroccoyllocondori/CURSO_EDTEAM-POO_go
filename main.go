package main

import "fmt"

// creacion de la  clase persona
type Persona struct {
	Nombre string
	Edad   uint
}

func NuevaPersona(nombre string, edad uint) Persona {
	return Persona{nombre, edad}
}

func (referncia Persona) Saludo() {
	fmt.Println("hola a todos")
}

// creacion de la estructura empleado a partir de persona
type Empleado struct {
	Persona // embebemos  la clase persona y automaticamente  hereda todo los atributos y metodos
	// el receptor de esos metodos sera el tipo interno
	sueldo float64
}

func NuevoEmpleado(nombre string, edad uint, sueldo float64) Empleado {
	return Empleado{NuevaPersona(nombre, edad), sueldo}
}
func (instancia Empleado) Honorario() {
	fmt.Println(instancia.sueldo * 0.9)
}

func main() {

	empleado := NuevoEmpleado("juan", 25, 960)
	fmt.Println(empleado.Nombre, empleado.Edad)
	empleado.Saludo()
	empleado.Honorario()

}
