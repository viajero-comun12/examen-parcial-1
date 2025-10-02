package main

import "fmt"

func main() {
	var num int
	fmt.Println("ingrese un numero entero positivo")
	fmt.Scan(&num)
	for i := 1; i <= num; i++ {

		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}

	}
	switch {

	case num >= 1 && num <= 10:
		fmt.Println("numero pequeño")
	case num >= 11 && num <= 100:
		fmt.Println("numero mediano")
	case num > 100:
		fmt.Println("Numero Grande")
	}

}
