package main

import "fmt"

func say1(text string, c chan<- string) {
	c <- text
}

func main_chanels() {
	c := make(chan string, 1)

	fmt.Println("Hello")

	go say1("Bye", c)

	fmt.Println(<-c)
}
