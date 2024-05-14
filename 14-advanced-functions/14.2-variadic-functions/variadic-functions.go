package main

import "fmt"

func soma(numeros ...int) int {
	fmt.Println(numeros)
	total := 0
	for _, numero := range numeros {
		total += numero
	}

	return total
}

func escrever(texto string, numeros ...int) {
	for _, numero := range numeros {
		fmt.Println(texto, numero)
	}
}

func main() {
	total := soma(1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 11, 2, 22, 3455)
	fmt.Println(total)
	escrever("Olá", 1, 2, 4, 5, 6, 7, 8, 9, 9)
}
