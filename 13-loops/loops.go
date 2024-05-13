package main

import (
	"fmt"
	"time"
)

func main() {
	type usuarioStruct struct {
		nome      string
		sobrenome string
	}

	i := 0
	for i < 10 {
		fmt.Println("add", i)
		time.Sleep(time.Second)
		i++
	}

	for j := 0; j < 10; j++ {
		fmt.Println(j)
		time.Sleep(time.Second)
	}

	nomes := []string{
		"João", "Davi", "Lucas",
	}

	for indice, nome := range nomes {
		fmt.Println(indice, nome)
	}

	for _, nome := range nomes {
		fmt.Println(nome)
	}

	for _, letra := range "letras" {
		fmt.Println(string(letra))
	}

	usuario := map[string]string{
		"nome":      "Carlos",
		"sobrenome": "Silva",
	}

	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}

	// NAO PODE FAZER FOR EM STRUCTS
	// usuario2 := usuarioStruct{"joao", "silva"}

	// for chave, valor := range usuario2 {
	// 	fmt.Println(chave, valor)
	// }
}
