package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Olá mundo!"))
}
func main() {
	// HTTP É UM PROTOCOLO CLIENTE SERVIDOR
	// REQUEST X RESPONSE

	http.HandleFunc("/home", home)

	log.Fatal(http.ListenAndServe(":5000", nil))

}
