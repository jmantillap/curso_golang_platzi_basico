package main

import (
	"fmt"
	"sync"
	"time"
)

func say(text string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(text)
}

func main_gorutines() {

	var wg sync.WaitGroup

	fmt.Println("Hello")

	wg.Add(1)

	go say("Word", &wg)

	wg.Wait()

	go func(text string) {
		fmt.Println(text)
	}("Adios")

	time.Sleep(time.Second * 1)
}
