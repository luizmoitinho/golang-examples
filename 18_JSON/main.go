package main

import (
	"bytes"
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
	fmt.Printf("Struct => JSON\n\n")
	c1 := cachorro{"Jack", "Husky", 2}
	//Converte map ou struct para json => marshal
	cachorroJSON, err := json.Marshal(c1)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(cachorroJSON)
	fmt.Println(bytes.NewBuffer(cachorroJSON))

	fmt.Printf("\n\nMap => JSON")
	c2 := map[string]string{
		"nome":  "Toby",
		"raca":  "Poodle",
		"idade": "2",
	}

	cachorroJSON, err = json.Marshal(c2)

	fmt.Println(cachorroJSON)
	fmt.Println(bytes.NewBuffer(cachorroJSON))

}
