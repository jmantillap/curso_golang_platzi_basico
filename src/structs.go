package main

import "fmt"

type car struct {
	brand   string
	year    int
	seating int
	color   string
	owner   string
}

type Person struct {
	name        string
	last        string
	phoneNumber uint32
}

func (p Person) getName() string {
	return p.name
}

func main_struct() {
	myCar := car{brand: "Toyota", year: 2018, seating: 10, color: "Rojo", owner: "Eliaz Bobadilla"}
	fmt.Println("Los Datos de mi auto son:", myCar)

	//
	var otherCar car
	otherCar.brand = "Ferrari"
	fmt.Println(otherCar)

	daniel := Person{
		name:        "Daniel",
		last:        "Franc",
		phoneNumber: 4061022,
	}

	println(daniel.last)
	daniel.last = "Franco"
	println(daniel.last)
	println(daniel.getName())
}
