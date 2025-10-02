package agenda

import "fmt"

var Agenda = map[string]int{}

func Agregar(nombre string, numero int) {
	Agenda[nombre] = numero
	fmt.Println("se agrego con exito")
}

func Buscar(nombre string) {
	numero, existe := Agenda[nombre]
	if existe {
		fmt.Println("si existe, nombre y numero: ", nombre, numero)
	} else {
		fmt.Println("no existe")
	}
}
func Eliminar(nombre string) {
	_, existe := Agenda[nombre]
	if existe {
		delete(Agenda, nombre)
		fmt.Println("se elimino")
	} else {
		fmt.Println("no se pudo eliminar no existe")
	}
}
func Listar() {
	for nombre, numero := range Agenda {
		fmt.Println(nombre, numero)
	}
}
