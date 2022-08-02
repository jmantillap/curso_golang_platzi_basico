package main

import "fmt"

func main_switch() {
	var modulo int = 6 % 2
	switch modulo {
	case 0:
		fmt.Println("Es un número par.")
	default:
		fmt.Println("Es un número impar.")
	}

	switch modulo1 := 6 % 2; modulo1 {
	case 0:
		fmt.Println("Es un número par.")
	default:
		fmt.Println("Es un número impar.")
	}

	var value int = -98
	switch {
	case value > 100:
		fmt.Println("Es un número mayor a 100.")
	case value < 0:
		fmt.Println("Es un número menor a 0.")
	default:
		fmt.Println("No cumple ninguna condición.")
	}

	whaType(3.14)
	whaType(4134)
	whaType("Go")

}

func whaType(value interface{}) {

	switch v := value.(type) {
	case float64:
		fmt.Printf("%g is float. \n", v)
	case int:
		fmt.Printf("%d is integer. \n", v)
	default:
		fmt.Printf("%v is %T. \n", v, v)
	}
}
