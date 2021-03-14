package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type cachorro struct {
	Nome  string `json: "nome"`
	Raca  string `json: "raca"`
	Idade int    `json: "idade"`
}

func main() {
	fmt.Printf("Convertendo JSON => Struct\n\n")
	c := cachorro{}
	cachorroJSON := `{"idade":2,"nome":"Toby","raca":"Poodle"}`

	err := json.Unmarshal([]byte(cachorroJSON), &c)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(c)

}
