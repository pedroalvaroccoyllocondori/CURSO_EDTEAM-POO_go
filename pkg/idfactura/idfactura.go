package idfactura

type Item struct {
	id       uint
	producto string
	valor    float64
}

func Nuevo(id uint, producto string, valor float64) Item {
	return Item{id, producto, valor} // se inicialializa los valores
}

type Itemsss []Item // nuevo tipo de estructura en base a items

func NuevoItems(items ...Item) Itemsss {
	var factura Itemsss
	for _, item := range items {
		factura = append(factura, item)
	}
	return factura
}

func (referencia Itemsss) Total() float64 {
	var total float64
	for _, item := range referencia {
		total += item.valor
	}
	return total
}
