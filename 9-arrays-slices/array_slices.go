package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Array and Slices")

	var array1 [5]int
	array1[0] = 1

	fmt.Println(array1)

	array2 := [5]string{"Posição1", "Posição2", "Posição3", "Posição4", "Posição5"}
	fmt.Println(array2)

	array3 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(array3)

	slice1 := []int{5, 4, 3, 2, 1}

	fmt.Println(reflect.TypeOf(slice1))
	fmt.Println(reflect.TypeOf(array1))

	fmt.Println(slice1)

	slice1 = append(slice1, 0)

	fmt.Println(slice1)

	slice2 := array2[1:3]

	fmt.Println(slice2)

	array2[1] = "Posição-Alterada"

	fmt.Println(slice2)

	// Arrays internos
	fmt.Println("-----------------------")
	slice3 := make([]float32, 10, 11)
	fmt.Println(slice3)

	slice3 = append(slice3, 5)
	slice3 = append(slice3, 5)
	fmt.Println(slice3)
	fmt.Println(len(slice3))
	fmt.Println(cap(slice3))

	slice4 := make([]float32, 5)
	fmt.Println(slice4)
	slice4 = append(slice4, 10)
	fmt.Println(len(slice4))
	fmt.Println(cap(slice4))

}
