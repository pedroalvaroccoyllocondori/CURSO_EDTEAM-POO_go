package idfactura

type Item struct {
	id       uint
	producto string
	valor    float64
}

func Nuevo(id uint, producto string, valor float64) Item {
	return Item{id, producto, valor}
}
func (instancia Item) Valor() float64 {
	return instancia.valor
}
