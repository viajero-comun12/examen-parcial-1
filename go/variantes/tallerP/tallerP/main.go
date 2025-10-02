package main

import (
	"fmt"
	taller "tallerP/packete"
)

func main() {
	var dolares float64
	var moneda int
	fmt.Println("ingrese su cantidad en dolares")
	fmt.Scan(&dolares)
	fmt.Println("elija su opcion 1:eruos, 2 lb, 3 wones, 4btc")
	fmt.Scan(&moneda)
	taller.Monedas(moneda, dolares)

	var frase string
	fmt.Println("ingrese su frase ")
	fmt.Scan(&frase)
	taller.Contador(frase)

}
