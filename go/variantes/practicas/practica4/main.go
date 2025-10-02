package main

import "fmt"

func main() {
	var num int
	var resp int
	fmt.Println("ingrese un numero entero positivo")
	fmt.Scan(&num)
	for i := 1; i <= 10; i++ {
		resp = num * i
		fmt.Println(num, "*", i, "=", resp)
	}
	if num%2 == 0 {
		fmt.Println("par")
	} else {
		fmt.Println("impar")
	}
	switch {

	case num >= 1 && num <= 5:
		fmt.Println("numero pequeño")
	case num >= 6 && num <= 10:
		fmt.Println("numero mediano")
	case num > 10:
		fmt.Println("Numero Grande")
	}

}
