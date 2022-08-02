package main

import "fmt"

type figuras2D interface {
	area() float64
}

type cuadrado struct {
	base float64
}

type rectangulo struct {
	base   float64
	altura float64
}

func (c cuadrado) area() float64 {
	return c.base * c.base
}

func (r rectangulo) area() float64 {
	return r.base * r.altura
}

func (c cuadrado) String() string {
	return fmt.Sprintf("Cuadrado de lado = %f", c.base)
}
func (r rectangulo) String() string {
	return fmt.Sprintf("Rectangulo Base = %f y Altura = %f", r.base, r.altura)
}

func calcular(f figuras2D) {
	fmt.Println("Area=", f.area())
}

func main_interfaces() {

	myCuadrado := cuadrado{base: 5}
	myRectangulo := rectangulo{base: 2, altura: 10}

	calcular(myCuadrado)
	calcular(myRectangulo)

	//lista de interfaces
	myInterface := []interface{}{"Hola", 12, 4.90}

	fmt.Println(myInterface...)

}
