package main

import (
	"fmt"
	"math"
)

func main() {
	//Declaración de constantes
	const pi float64 = 3.14
	const pi2 = 3.1415

	fmt.Println("pi:", pi)
	fmt.Println("pi2:", pi2)

	//Declaracion de variables enteras
	base := 12          //Primera forma
	var altura int = 14 //Segunda forma
	var area int        //Go no compila si las variables no son usadas

	fmt.Println(base, altura, area)

	//Zero values
	//Go asigna vaalores a variables vacías
	var a int
	var b float64
	var c string
	var d bool

	fmt.Println(a, b, c, d)

	//Ejercicio
	//Calcular el áera del cuadrado
	const baseCuadrado = 10
	areaCuadrado := baseCuadrado * baseCuadrado

	fmt.Println("El área del cuadrado es:", areaCuadrado)

	x := 10
	y := 50

	//Suma
	result := x + y
	fmt.Println("Suma:", result)

	//Resta
	result = y - x
	fmt.Println("Resta:", result)

	//Multiplicacion
	result = x * y
	fmt.Println("Multiplicació:", result)

	//División
	result = y / x
	fmt.Println("División:", result)

	//Modulo
	result = y % x
	fmt.Println("Modulo:", result)

	//Incremental
	x++
	fmt.Println("Incremetal:", x)

	//Decremental
	x--
	fmt.Println("Decremental:", x)

	//Retos
	//Area de un rectangulo, trapecio y de un circulo
	//Rectangulo
	ladoA := 4
	ladoB := 3
	areaRectangulo := ladoA * ladoB
	fmt.Println("Área de un rectángulo:", areaRectangulo)

	//Trapecio
	a = 5
	b = 8
	areaTrapecio := (5 + 8) / 2
	fmt.Println("Área de un trapecio:", areaTrapecio)

	//Circulo
	diametro := 5
	pi3 := math.Pi
	diametroCuadrado := math.Pow(float64(diametro), 2)
	areaCirculo := (pi3 * diametroCuadrado) / 4
	fmt.Println("Área de un circulo:", math.Round(areaCirculo*100)/100)
}
