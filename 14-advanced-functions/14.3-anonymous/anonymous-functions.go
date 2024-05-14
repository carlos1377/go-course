package main

import "fmt"

func main() {
	retorno := func(text string) string {
		return fmt.Sprintf("Recebido -> %s", text)
	}("olá")

	fmt.Println(retorno)
}
