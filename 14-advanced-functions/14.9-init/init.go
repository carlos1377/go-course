package main

import "fmt"

var n int

func main() {
	fmt.Println("func main", n)

}

// sempre acontece antes da main, pode ter uma por arquivo
func init() {
	n = 10
	fmt.Println("init func")
}
