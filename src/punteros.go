package main

import "fmt"

type pc struct {
	brand string
	disk  int
	ram   int
}

func (myPC pc) ping() {
	fmt.Println(myPC.brand, "pong")
}

func (myPC *pc) duplicateRam() {
	myPC.ram = myPC.ram * 2
}

func (miPc *pc) toString() string {
	return fmt.Sprintf("Tengo %dGB de RAM, %dGB de disco y marca %s", miPc.ram, miPc.disk, miPc.brand)

}

func main_puneros() {
	fmt.Println("")
	fmt.Println("==================")
	fmt.Println("Structs y Punteros")
	fmt.Println("=================")
	fmt.Println()

	a := 50
	b := &a
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(*b)

	*b = 100
	fmt.Println(a)

	myPC := pc{ram: 16, disk: 200, brand: "MSI"}

	fmt.Println(myPC)

	myPC.ping()

	fmt.Println(myPC)
	myPC.duplicateRam()

	fmt.Println(myPC)
	myPC.duplicateRam()

	fmt.Println(myPC.toString())

}
