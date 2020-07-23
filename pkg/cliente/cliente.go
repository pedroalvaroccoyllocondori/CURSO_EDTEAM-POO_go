package cliente

type Cliente struct {
	nombre    string
	direccion string
	telefono  string
}

func Nuevo(nombre, direccion, telefono string) Cliente {
	return Cliente{nombre, direccion, telefono}
}
