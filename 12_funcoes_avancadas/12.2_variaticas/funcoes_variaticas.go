package main

import "fmt"

func soma(numeros ...int) int {
	var resultado int
	for _, valor := range numeros {
		resultado += valor
	}
	return resultado

}

func escrever(texto string, numeros ...int) {
	for _, valor := range numeros {
		fmt.Println(texto, valor)
	}
}

func main() {
	fmt.Println("Funções variáticas")
	fmt.Println("Resultado", soma(10, 30, 63, 99))
	fmt.Println("Mais de um parametro")
	escrever("golang", 10, 30, 63, 99)
}
