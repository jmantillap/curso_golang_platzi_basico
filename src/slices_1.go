package main

import (
	"fmt"
	"strings"
)

func main_slices() {

	slice := []string{"hola", "que", "hace"}

	for i, valor := range slice {
		fmt.Println(i, valor)
	}

	for _, valor := range slice {
		fmt.Println(valor)
	}

	for i := range slice {
		fmt.Println(i)
	}

	isPalindrome("amor a Roma")

}

func isPalindrome(text string) {
	text = strings.ToLower(text)

	var textReverse string

	for i := len(text) - 1; i >= 0; i-- {
		textReverse += string(text[i])
	}

	if text == textReverse {
		fmt.Println("Es palindromo")
	} else {
		fmt.Println("No es palindromo")
	}
}
