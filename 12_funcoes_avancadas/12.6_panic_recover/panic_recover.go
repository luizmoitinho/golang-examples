package main

import "fmt"

func recuperarExecucao() {
	if r := recover(); r != nil {
		fmt.Println("Execução do programa foi recuperada.")
	}

}

func divisao(n1, n2 float64) float64 {
	defer recuperarExecucao()
	if n2 == 0 {
		panic("O divisor não pode ser igual a 0!")
	}

	return n1 / n2
}

func main() {
	fmt.Println("# Panic e Defer")
	fmt.Println(divisao(10, 2))
}
