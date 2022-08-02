package main

import "fmt"

func say2(text string, c chan<- string) {
	c <- text
}

func main() {
	c := make(chan string, 2)
	c <- "Mensaje 1"
	c <- "Mensaje 2"

	fmt.Println(len(c), cap(c))
	//go say1("Bye", c)
	//fmt.Println(<-c)

	//range y close
	close(c)
	//c <- "Mensaje 3"
	for message := range c {
		fmt.Println(message)
	}

	//select
	email1 := make(chan string)
	email2 := make(chan string)

	fmt.Println("------SELECT-------")
	go message("mensaje 3", email1)
	go message("mensaje 4", email2)

	// tener presente la cantidad de channels a manejar para el ciclo, en este ejemplo 2
	for i := 0; i < 2; i++ {
		select {
		case m1 := <-email1:
			fmt.Println("Email recibido de email 1", m1)
		case m2 := <-email2:
			fmt.Println("Email recibido de email 2", m2)
		}
	}

}

//del ejemplo message SELECT
func message(text string, c chan string) {
	c <- text
}
