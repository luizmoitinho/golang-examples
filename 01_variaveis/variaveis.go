package main

import "fmt"

func main() {
	var nome string = "Luiz Carlos"
	sobrenome := "Costa Moitinho"

	fmt.Println("Nome: ", nome)
	fmt.Println("Sobrenome: ", sobrenome)

	fmt.Println("================================")

	var (
		nomeCompleto string
		idade        int = 22
	)
	nomeCompleto = nome + " " + sobrenome
	fmt.Println("Nome Completo: ", nomeCompleto)
	fmt.Println("Idade: ", idade)

	fmt.Println("================================")

	const PI = 3.14
	fmt.Println("Valor de PI = ", PI)

	fmt.Println("================================")
	fmt.Println("Invertendo valores")

	num1, num2 := 6, 7
	fmt.Println("num1 = ", num1, "| num2 = ", num2)

	num1, num2 = num2, num1
	fmt.Println("num1 = ", num1, "| num2 = ", num2)

}
