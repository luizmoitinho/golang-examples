package main

import "fmt"

func closure() func() {
	texto := "Dentro da função closure"

	display := func() {
		fmt.Println(texto)
	}

	return display

}

func main() {

	texto := "Dentro da função Main"
	fmt.Println(texto)

	funcAux := closure()

	funcAux()

}
