package main

import (
	"fmt"
	"log"
	"strconv"
)

func main_if() {
	valor1 := 1
	valor2 := 2
	if valor1 == 1 {
		fmt.Println("Es 1")
	} else {
		fmt.Println("No Es 1")
	}

	if valor1 == 1 && valor2 == 2 {
		fmt.Println("Es verdad AND")
	}

	if valor1 == 1 || valor2 == 3 {
		fmt.Println("Es verdad OR")
	}

	//Convertir texto a numero
	value, err := strconv.Atoi("53")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("value:", value)

	if isPair(6) {
		fmt.Println("Number is pair")
	} else {
		fmt.Println("Number is odd")
	}
	if isValidUser("Alpha5", "MyPassword") {
		fmt.Println("Credentials are valid")
	} else {
		fmt.Println("Credentials aren't valid")
	}
}

func isPair(num int) bool {
	return num%2 == 0
}

func isValidUser(userName, pass string) bool {
	return userName == "Alpha" && pass == "MyPassword"
}
