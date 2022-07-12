package main

import "fmt"

func normalFunction(message string) {
	fmt.Println(message)
}

func tripeArgument(a, b int, c string) {
	fmt.Println(a, b, c)
}

func returnValue(a int) int {
	return a * 2
}

func doubleReturn(a int) (c, d int) {
	return a, a * 2
}

func areaCuadrado(base int) (area int) {
	return base * base
}

func areaRectangulo(base, height int) (area int) {
	return base * height
}

func areaTrapecio(base1, base2, height int) (area int) {
	return ((base1 + base2) / 2) * height
}

func areaCirculo(diametro int, PI float64) (area float64) {
	radio := diametro / 2
	return PI * ((float64(radio)) * (float64(radio)))
}

func main() {

	normalFunction("Hola mundo")
	tripeArgument(1, 2, "hola")

	value := returnValue(2)
	fmt.Println("Value:", value)

	value1, value2 := doubleReturn(2)
	fmt.Println("value1 y value2:", value1, value2)

	const PI float64 = 3.1415926535
	x := 10
	y := 50
	z := 16

	// Area cuadrado
	fmt.Println("Area cuadrado:", areaCuadrado(x))

	// Area Rectángulo
	fmt.Println("Area Rectángulo:", areaRectangulo(x, y))

	// Area Trapecio
	fmt.Println("Area Trapecio:", areaTrapecio(x, y, z))

	// Area Círculo
	fmt.Println("Area Circulo:", areaCirculo(x, PI))
	fmt.Println("Area Circulo:", areaCirculo(y, PI))
	fmt.Println("Area Circulo:", areaCirculo(z, PI))
}
