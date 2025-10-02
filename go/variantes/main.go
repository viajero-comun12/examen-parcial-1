package main


import "fmt"

func main() {
	var edad int
	fmt.Println("ingrese su edad")
	fmt.Scan(&edad)
	if edad >= 18 {
		fmt.Println("eres mayor de edad")
	} else {
		fmt.Println("eresmenor de edad")
	}
	//switch//
	var dia string
	fmt.Println("ingrese el dia de hoy ")
	fmt.Scan(&dia)

	switch dia {
	case "lunes":
		fmt.Println("primer dia de la semana")
	case "miercoles":
		fmt.Println("mitad de semana")
	case "viernes":
		fmt.Println("fin de semana laboral")
	case "domingo":
		fmt.Println("fin de semana")
	default:
		fmt.Println("dia no especificado dentro de lo programado")
	}
	//bucles//
	var i int
	for i = 0; i < 5; i++ {
		fmt.Println("i =", i)
	}
	//for para todo while = for solo con condicion
	for i < 5 {
		fmt.Println("si")
	}
	//while
	for {
		fmt.Println("sicopara")
		i++
		if i > 5 {
			break
		}
	}
	// arreglos
	var cant int
	fmt.Println("ingrese el tamño de numero para el arreglo")
	fmt.Scan(&cant)
	num := make([]int, cant)

	for i := 0; i < cant; i++ {
		fmt.Println("ingrese su numero")
		fmt.Scan(&num[i])
	}
	fmt.Println("los numeros ingresados son")
	for i, v := range num {
		fmt.Println("indice", i, "valor ", v)

	}
}
