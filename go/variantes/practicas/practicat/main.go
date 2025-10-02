package main

import (
	"fmt"
	"practicat/agenda"
)

func main() {
	var opcion int
	estadisticas := [3]int{0, 0, 0}
	for {
		fmt.Println("menu aceptado")
		fmt.Println("1. Agregar contacto")
		fmt.Println("2. Buscar contacto")
		fmt.Println("3. Eliminar contacto")
		fmt.Println("4. Listar contactos")
		fmt.Println("5. Ver estadísticas")
		fmt.Println("6. Salir")
		fmt.Print("Seleccione una opción: ")
		fmt.Scan(&opcion)

		switch opcion {
		case 1:
			var nombre string
			var numero int
			fmt.Println("ingrese el nombre")
			fmt.Scan(&nombre)
			fmt.Println("ingrese el numero")
			fmt.Scan(&numero)
			agenda.Agregar(nombre, numero)
			estadisticas[0]++
		case 2:
			var nombre string
			fmt.Println("ingrtese el nombre a buscar")
			fmt.Scan(&nombre)
			agenda.Buscar(nombre)
			estadisticas[2]++
		case 3:
			var nombre string
			fmt.Println("ingrtese el nombre a eliminar")
			fmt.Scan(&nombre)
			agenda.Eliminar(nombre)
			estadisticas[1]++
		case 4:
			agenda.Listar()
		case 5:
			fmt.Println(estadisticas[0], estadisticas[1], estadisticas[2])
		case 6:
			fmt.Println("chao")
			return
		default:
			fmt.Println("es estupido")
		}
	}
}
