package main

import "fmt"

func main() {
	fmt.Println("Maps")

	usuario := map[int]string{
		0: "Luiz Carlos",
		1: "Carlos Alberto",
	}
	fmt.Println(usuario[0])
	fmt.Printf("\nMaps aninhados\n\n")
	usuario2 := map[string]map[string]string{
		"nome": {
			"primeiro": "Luiz Carlos",
			"ultimo":   "Costa Moitinho",
		},
		"curso": {
			"nome":         "Sistemas de Informação",
			"departamento": "DSI",
			"campus":       "Prof. Alberto Carvalho - Itabaiana",
			"universidade": "UFS",
		},
	}

	fmt.Println(usuario2)

	fmt.Printf("\nDeletando uma chave: 'nome' \n\n")
	delete(usuario2, "nome")
	fmt.Println(usuario2)

	usuario2["disciplinas"] = map[string]string{
		"001": "Programação I",
		"002": "Cálculo I",
		"003": "Introdução a Administração",
	}

	fmt.Printf("\n Adicionando uma chave: 'disciplinas' \n\n")
	fmt.Println(usuario2)

}
