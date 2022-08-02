package main

import "fmt"

func main_map() {
	m := make(map[string]int)

	m["Jose"] = 14
	m["Pepito"] = 20

	fmt.Println(m)

	// Recorrer un map
	for i, v := range m {
		fmt.Printf("%s tiene %d años\n", i, v)
	}
	// Encontrar un valor

	value := m["Jose"]
	fmt.Println(value)

	value, ok := m["Jose"]
	fmt.Println(value, ok)
}
