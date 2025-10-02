package main

import "fmt"

func main() {
	var num int
	var resp int
	var residuo_total int

	fmt.Println("ingrese un numero entero positivo")
	fmt.Scan(&num)
	for i := 0; num > 0; i++ {
		residuo := num % 10
		residuo_total += residuo

		num /= 10
		resp++
		fmt.Println(resp)

	}

	fmt.Println("la suma es de ", residuo_total)

	switch {

	case resp == 1:
		fmt.Println("numero de una cifra")
	case resp == 2:
		fmt.Println("numero de dos cifras")
	case resp == 3:
		fmt.Println("numero de 3 cifras")
	case resp > 3:
		fmt.Println("numero de varias cifras")
	}
}
