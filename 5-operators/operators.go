package main

import "fmt"

func main() {
	// ARITMETICOS
	soma := 1 + 2
	subtracao := 2 - 1
	divisao := 10 / 2
	multiplicacao := 5 * 5
	restoDaDivisao := 10 % 3

	fmt.Println(soma, subtracao, divisao, multiplicacao, restoDaDivisao)

	var numero1 int16 = 10
	var numero2 int32 = 25

	soma2 := numero1 + int16(numero2)

	fmt.Println(soma2)

	// ATRIBUIÇÃO
	var variavel string = "String"
	variavel2 := "string2"

	fmt.Println(variavel, variavel2)

	// RELACIONAIS
	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 < 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 != 2)
	fmt.Println(1 == 2)

	// LÓGICOS
	fmt.Println(1 != 3 && 2 != 2)
	fmt.Println(1 != 3 || 2 != 2)
	fmt.Println(!true)

	// OPERADORES UNÁRIOS
	numero := 10
	numero++
	numero += 15
	fmt.Println(numero)
	numero--
	fmt.Println(numero)
	numero *= 3
	fmt.Println(numero)
	numero /= 2
	fmt.Println(numero)

	// GO NAO TEM TERNARIOS

	var texto string
	if numero > 5 {
		texto = "Maior que 5"
	} else {
		texto = "Menor que 5"
	}
	fmt.Println(texto)
}
