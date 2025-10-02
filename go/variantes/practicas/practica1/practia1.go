package main

import "fmt"

func main() {
	var cant int

	fmt.Println("ingrese cuantos numeros quiere ingresar")
	fmt.Scan(&cant)
	num := make([]int, cant)

	for i := 0; i < cant; i++ {
		fmt.Println("ingrese los numeros ")
		fmt.Scan(&num[i])
	}
	for i, v := range num {
		if v%2 == 0 {
			fmt.Println("es un numero par ", i, v)
		} else {
			fmt.Println("es un numero impar", i, v)
		}
	}
}
