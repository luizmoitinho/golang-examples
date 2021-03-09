package main

import "fmt"

func main() {
	fmt.Println("Estruturas de controle")

	num := 10

	if num%2 == 0 {
		fmt.Println("Par")
	} else {
		fmt.Println("Impar")
	}

	if numAux := num; numAux > 0 {
		fmt.Println("Numero é maior que zero")
	} else {
		fmt.Println("Numero é menor que zero")
	}
}
