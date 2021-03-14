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
	c1 := cachorro{"Jack", "Husky", 2}
	//Converte map ou struct para json => marshal
	cachorroJSON, err := json.Marshal(c1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(cachorroJSON)
}
