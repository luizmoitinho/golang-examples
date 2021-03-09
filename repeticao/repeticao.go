package main

import (
	"fmt"
)

func main() {
	fmt.Println("Estrutura de Repetição")

	i := 0
	for i < 10 {
		fmt.Printf("%2d", i)
		//time.Sleep(time.Second)
		i++
	}

	fmt.Printf("\nResultado: %d\n", i)

	fmt.Printf("\nInterando em J junto com for\n")
	for j := 0; j < 10; j++ {
		fmt.Printf("%2d", j)
		//time.Sleep(time.Second)
	}

	fmt.Printf("\n\nFor com range")
	nome := [...]string{"luiz", "carlos", "costa", "moitinho"}
	for indice, valor := range nome {
		fmt.Printf("\nnome[%d] = %s", indice, valor)
	}

	fmt.Printf("\n\nFor com range: descartando o indice\n")
	for _, valor := range nome {
		fmt.Printf("%s\n", valor)
	}

	fmt.Printf("\n\niterando sobre uma string\n")
	for indice, letra := range "luiz carlos" {
		fmt.Printf(" [%d,%s] ", indice, string(letra))

	}

	fmt.Printf("\n\niterando sobre um map\n")
	usuario := map[string]string{
		"nome":      "Luiz carlos",
		"sobrenome": "costa moitinho",
	}

	for indice, valor := range usuario {
		fmt.Println(indice, ":", valor)
	}
}
