package main

import (
	"fmt"
	"time"
)

func main() {
	canal := make(chan string)

	go escrever("olá mundo", canal)

	for {
		mensagem, aberto := <-canal
		if !aberto {
			break
		}
		fmt.Println(mensagem)
	}

	// Outra forma de fazer o código acima
	// for mensagem := range canal {
	// 	fmt.Println(mensagem)
	// }

	fmt.Println("Fim do programa")
}

func escrever(texto string, canal chan string) {
	for i := 0; i < 5; i++ {
		canal <- texto
		time.Sleep(time.Second)
	}

	close(canal)
}
