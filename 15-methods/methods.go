package main

import "fmt"

type usuario struct {
	nome  string
	idade uint8
}

func (u usuario) salvar() {
	fmt.Printf("Salvado %s\n", u.nome)
}

func (u usuario) maiorDeIdade() bool {
	return u.idade >= 18
}

func (u *usuario) fazerAniversario() {
	u.idade++
}

func main() {
	usuario1 := usuario{"Usuario 1", 18}

	fmt.Println(usuario1)
	usuario1.salvar()

	usuario2 := usuario{"Carlos", 15}

	usuario2.salvar()

	maior := usuario2.maiorDeIdade()
	fmt.Println(maior)

	usuario2.fazerAniversario()
	fmt.Println(usuario2.idade)
}
