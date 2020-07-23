package factura

import "github.com/pedroalvaroccoyllocondori/POO_go/tree/master/pkg/cliente"

type factura struct {
	pais   string
	ciudad string
	total  float64
	client cliente.Cliente
}
