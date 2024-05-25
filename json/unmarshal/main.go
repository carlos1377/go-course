package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type cachorro struct {
	Nome  string `json:"nome"`
	Raca  string `json:"raca"`
	Idade uint   `json:"idade"`
}

func main() {
	cachorroEmJSON := `{"nome":"Rex","raca":"Dálmata","idade":3}`

	var c cachorro

	if err := json.Unmarshal([]byte(cachorroEmJSON), &c); err != nil {
		log.Fatal(err)
	}
	fmt.Println(c)

	cachorro2EmJSON := `{"nome":"Tob","raca":"Poodle"}`

	c2 := make(map[string]string)

	if err := json.Unmarshal([]byte(cachorro2EmJSON), &c2); err != nil {
		log.Fatal(err)
	}

	fmt.Println(c2)
}
