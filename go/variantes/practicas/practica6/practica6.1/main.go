package main

import "fmt"

func agregar(mapa1 map[string]int) {
	var nombre string
	var numero int
	fmt.Println("ingrese el nombre y el numero de telefono a agregar")
	fmt.Scan(&nombre, &numero)
	mapa1[nombre] = numero

}
func buscar(mapa1 map[string]int) {
	var busc string
	fmt.Println("ingrese que desea buscar en nombre")
	fmt.Scan(&busc)

	for nombre, numero := range mapa1 {
		if busc == nombre {
			fmt.Println("el contacto es", nombre, "y su numero es", numero)
		} else {
			fmt.Println("no se encontro")
		}
	}

}
func elim(mapa1 map[string]int) {
	var nombre string
	fmt.Println("ingrese el nombre a eliminar")
	fmt.Scan(&nombre)
	_, existe := mapa1[nombre]
	if existe {
		fmt.Println("el contacto es", nombre, "se a elimnado")
		delete(mapa1, nombre)
	} else {
		fmt.Println("no se encontro")
	}
}

func listar(mapa1 map[string]int) {
	for nombre, numero := range mapa1 {
		fmt.Println("el contacto es", nombre, "y su numero es", numero)
	}
}
func main() {
	mapa1 := map[string]int{}

	var selec int = 1
	var op int
	for i := 0; selec == 1; i++ {
		fmt.Println("ingrese la opcion a selecionar")
		fmt.Println("1agregar")
		fmt.Println("2buscar")
		fmt.Println("3eliminar")
		fmt.Println("4listar")
		fmt.Println("5salir")
		fmt.Scan(&op)
		switch {
		case op == 1:
			agregar(mapa1)
		case op == 2:
			buscar(mapa1)
		case op == 3:
			elim(mapa1)
		case op == 4:
			listar(mapa1)
		case op == 5:
			return
		}
	}
}
