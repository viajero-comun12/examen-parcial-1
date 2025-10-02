package taller

import "fmt"

func Monedas(moneda int, dolares float64) {

	var resp float64
	var euros float64 = 0.86
	var lb float64 = 0.75
	switch moneda {
	case 1:
		resp = dolares * euros
		fmt.Println("en euros es ", resp)
	case 2:
		resp = dolares * lb
		fmt.Println("en libras es", resp)
	case 3:
		resp = dolares * 1409.20
		fmt.Println("en wones", resp)
	case 4:
		resp = dolares * 0.0000092
		fmt.Println("en bitcoins", resp)
	}
}

func Contador(cadena string) {
	vocales := []rune{'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U', 'Á', 'E', 'I', 'O', 'U', 'á', 'é', 'í', 'ó', 'ú'}
	contador := 0
	for _, caracter := range cadena {
		for _, vocal := range vocales {
			if caracter == vocal {
				contador++
				break
			}
		}
	}
	fmt.Println("el numero es de", contador)
}
