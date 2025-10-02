package main

import "fmt"

func main() {
	numeros := []int{}
	var num int
	var op int
	fmt.Println("ingrese la cantidad de numeros a ingresar")
	fmt.Scan(&num)
	for i := 0; i < num; i++ {
		fmt.Println("ingrese el numero de la opcion", i)
		fmt.Scan(&op)
		numeros = append(numeros, op)
	}

	par, inpar := parimpar(numeros)
	fmt.Println("ahora vera los slices")
	fmt.Println("numeros almacenados", numeros)
	fmt.Println("pares ", par)
	fmt.Println("inpares", inpar)
}

func parimpar(nums []int) ([]int, []int) {
	par := []int{}
	inpar := []int{}

	for _, numero := range nums {

		if numero%2 == 0 {
			par = append(par, numero)
		} else {
			inpar = append(inpar, numero)
		}
	}
	return par, inpar
}
