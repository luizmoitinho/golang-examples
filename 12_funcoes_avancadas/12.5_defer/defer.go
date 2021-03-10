package main

import "fmt"

func func1() {
	fmt.Println("Executando a func1")
}

func func2() {
	fmt.Println("Executando a func2")
}

func isAlunoAprovado(n1, n2 float32) bool {
	defer fmt.Println("Média calculada. Resultado será retornado!")
	fmt.Println("Entrando na função para verificar se o aluno está aprovado")

	if (n1+n2)/2 >= 6 {
		return true
	}
	return false
}

func main() {
	fmt.Printf("# Defer\n\n")

	fmt.Println("# Sem o Defer:")
	func1()
	func2()

	fmt.Println("")
	fmt.Println("# Com o Defer:")

	defer func1()
	func2()

	fmt.Println("")
	fmt.Println("# Situação do aluno:")
	fmt.Println(isAlunoAprovado(10, 2))
}
