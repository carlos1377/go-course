package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     uint8
	altura    uint8
}

type estudante struct {
	pessoa
	curso     string
	faculdade string
}

func main() {
	fmt.Println("herança")

	p1 := pessoa{"carlos", "banana", 14, 180}

	fmt.Println(p1)

	e1 := estudante{p1, "ADS", "FSG"}

	fmt.Println(e1)
}
