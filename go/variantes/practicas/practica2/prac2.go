package main

import "fmt"

func main() {
	var num1, num2, r1, r2, r3, r4 int
	var op string
	fmt.Println("ingrese 2 numero enteros")
	fmt.Scan(&num1, &num2)
	fmt.Println("que operacion desea realizar suma, resta multiplicacion , division")
	fmt.Scan(&op)
	switch op {
	case "suma":
		r1 = num1 + num2
		fmt.Println("la suma es de ", r1)
	case "resta":
		r2 = num1 - num2
		fmt.Println("la resta es de ", r2)
	case "multiplicacion":
		r3 = num1 * num2
		fmt.Print("la multiplicacion es de", r3)
	case "division":
		r4 = num1 / num2
		fmt.Println("la division es de ", r4)
	default:
		fmt.Print("su operacion no es valida")
	}

}
