package main

import "fmt"

type computer struct {
	brand string
	disk  int
	ram   int
}

func (miPc computer) String() string {
	return fmt.Sprintf("Tengo %dGB de RAM, %dGB de disco y una targeta grafica %s", miPc.ram, miPc.disk, miPc.brand)

}

func main_stringers() {

	myPc := computer{ram: 16, brand: "MSI", disk: 200}

	fmt.Println(myPc)
}
