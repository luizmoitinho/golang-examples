package main

import (
	"fmt"
	"modulo/auxiliar"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Escrevendo no arquivo main!!")
	auxiliar.Escrever()

	email := "luizcarlos_costam@hotmail.com"
	err := checkmail.ValidateFormat(email)

	if err != nil {
		fmt.Println("E-mail inválido:", email)
		return
	}
	fmt.Println("E-mail valido:", email)

}
