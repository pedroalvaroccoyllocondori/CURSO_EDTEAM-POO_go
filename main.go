package main

import (
	"fmt"
)

type MetodoPago interface {
	Pagar()
}

type Paypal struct{}

func (referencia Paypal) Pagar() { // implemtacion de la interfaz
	fmt.Println("pagado por paypal")
}

type Tarjeta struct{}

func (referencia Tarjeta) Pagar() { // implemtacion de la interfaz
	fmt.Println("pagado  con tarjeta")
}

type Efectivo struct{}

func (referencia Efectivo) Pagar() { // implemtacion de la interfaz
	fmt.Println("pagado por  efectivo")
}

func Cajero(tipopago uint) MetodoPago { // metodo que devuelve una interfaz
	switch tipopago {
	case 1:
		return Paypal{}
	case 2:
		return Tarjeta{}
	case 3:
		return Efectivo{}
	default:
		return nil

	}
}

func main() {

	var metodo uint // declaracion de la variable
	fmt.Println("digite el metodo de pago \n 1: paypal \n 2: tarjeta \n 3: efectivo")
	_, error := fmt.Scanln(&metodo)
	if error != nil { // si se inserta letra
		panic("debes diguitar un numero entre 1 y 3 y no letras")
	}
	if metodo > 3 { // si se insert numeros mayores a 3
		panic("debes diguitar un numero entre 1 y 3 y no mayores a ese rango")
	}

	metodopago := Cajero(metodo) //el metodo devuelve una interfaz
	metodopago.Pagar()

}
