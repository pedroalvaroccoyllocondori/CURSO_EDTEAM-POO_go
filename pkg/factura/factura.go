package factura

import (
	"github.com/pedroalvaroccoyllocondori/POO_go/pkg/cliente"
	"github.com/pedroalvaroccoyllocondori/POO_go/pkg/idfactura"
)

// estuctura de la factura
type Factura struct {
	pais   string
	ciudad string
	total  float64
	client cliente.Cliente
	item   idfactura.Itemsss // los corchetes me indica relacion de 1a  muchos
	//creando un slice de  id fatura
}

// funcion constructora de  la estructura (clase)
func Nuevo(pais, ciudad string, client cliente.Cliente, item idfactura.Itemsss) Factura {

	return Factura{
		pais:   pais,
		ciudad: ciudad,
		client: client,
		item:   item,
	}
}

//inicializador del atributo total

func (instancia *Factura) Establecertotal() { //* para actualizar el valor
	instancia.total = instancia.item.Total()
}
