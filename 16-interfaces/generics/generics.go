package main

import "fmt"

func generic(interf interface{}) {
	fmt.Println(interf)
}

func main() {
	generic("String")
	generic(1)
	generic(true)

	// bagunça
	mapa := map[interface{}]interface{}{
		1:            "String",
		float64(100): true,
		"String":     uint8(20),
	}

	fmt.Println(mapa)
}
