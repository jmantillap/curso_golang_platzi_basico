package main

import (
	mypackage "curso_golang_platzi/src/mypackage"
	"fmt"
)

func main_modificadores() {
	fmt.Println("Inicio")

	var myCar mypackage.CarPublic
	myCar.Brand = "Ferrari"
	myCar.Year = 2020
	//myCar.otroYear = 2020
	fmt.Println("myCar", myCar)

	//var myOtherCar mypackage.carPrivate
	//fmt.Println("myOtherCar", myOtherCar)
	mypackage.PrintMessage("Hola platzi")
	//mypackage.printMessage1("Hola platzi")
}
