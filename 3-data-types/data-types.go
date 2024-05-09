package main

import (
	"errors"
	"fmt"
)

func main() {
	var numero int64 = 10000000000
	fmt.Println(numero)

	var numero2 uint32 = 1000
	fmt.Println(numero2)

	// alias INT32 = RUNE
	var numero3 rune = 123456
	fmt.Println(numero3)

	// BYTE = UINT8
	var numero4 byte = 123
	fmt.Println(numero4)

	var numeroReal1 float32 = 123.45
	fmt.Println(numeroReal1)

	var numeroReal2 float64 = 123000000.45
	fmt.Println(numeroReal2)

	numeroReal3 := 12345.67
	fmt.Println(numeroReal3)

	var str string = "Texto"
	fmt.Println(str)

	str2 := "Texto2"
	fmt.Println(str2)

	char := 'B'
	fmt.Println(char)

	// retorna um valor vazio daquele type
	var blank int16 // string
	fmt.Println(blank)

	var booleano1 bool = true
	fmt.Println(booleano1)

	var err error
	fmt.Println(err)

	var err_new error = errors.New("internal error")
	fmt.Println(err_new)
}
