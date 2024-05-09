package main

import "fmt"

func somar(x int8, y int8) int8 {
	return x + y
}

func calculosMatematicos(x, y int8) (int8, int8) {
	soma := x + y
	subtracao := x - y

	return soma, subtracao
}

func main() {
	soma := somar(10, 20)
	fmt.Println(soma)

	var f = func(text string) string {
		fmt.Println(text)

		return text
	}

	resultado := f("text of f")
	fmt.Println(resultado)

	resultadoSoma, resultadoSubtracao := calculosMatematicos(10, 15)
	fmt.Println(resultadoSoma, resultadoSubtracao)

	resultadoSoma1, _ := calculosMatematicos(20, 15)
	fmt.Println(resultadoSoma1)
}
