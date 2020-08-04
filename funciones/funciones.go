package funciones

import (
	"fmt"
	"time"
)

func Saludar(nombre string) {
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	fmt.Println("hola", nombre)
}

func Despedirse(nombre string) {
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	fmt.Println("adios", nombre)
}
