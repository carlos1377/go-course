package main

import "fmt"

type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}

type endereco struct {
	logradouro string
	numero     uint8
}

func main() {
	var user usuario
	user.nome = "carlos"
	user.idade = 21
	fmt.Println(user)

	endereco := endereco{"bairro x", 55}

	user2 := usuario{"carlos", 24, endereco}
	fmt.Println(user2)

	user3 := usuario{idade: 21}
	fmt.Println(user3)

}
