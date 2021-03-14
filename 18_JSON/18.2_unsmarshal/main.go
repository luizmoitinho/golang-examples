package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type cachorro struct {
	ID    int    `json: -`
	Nome  string `json: "nome"`
	Raca  string `json: "raca"`
	Idade int    `json: "idade"`
}

func main() {
	fmt.Printf("Convertendo JSON => Struct\n\n")
	c := cachorro{}
	cachorroJSON := `{"idade":2,"nome":"Toby","raca":"Poodle"}`

	if err := json.Unmarshal([]byte(cachorroJSON), &c); err != nil {
		log.Fatal(err)
	}

	fmt.Println(c)

	fmt.Printf("\n\nConvertendo JSON => Map")
	cachorro2JSON := `{"nome":"Toby","raca":"Poodle"}`
	c2 := make(map[string]string)
	if err := json.Unmarshal([]byte(cachorro2JSON), &c2); err != nil {
		log.Fatal(err)
	}

	fmt.Println(c2)
}
