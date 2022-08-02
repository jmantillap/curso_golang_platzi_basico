package mypackage

import "fmt"

type car struct {
	brand string
	year  int
}

// CarPublic Car con acceso publico
type CarPublic struct {
	Brand    string
	Year     int
	otroYear int
}

// carPrivate
type carPrivate struct {
	brand string
	year  int
}

// PrintMessage imprimir el mensaje
func PrintMessage(text string) {
	fmt.Println(text)
}

// printMessage1
func printMessage1(text string) {
	fmt.Println(text)
}
