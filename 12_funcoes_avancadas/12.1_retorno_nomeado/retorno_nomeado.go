package main

import "fmt"

func calcMath(n1, n2 int) (soma int, subtracao int) {
	soma = n1 + n2
	subtracao = n1 - n2
	return
}

func main() {
	fmt.Printf("\nFunções com retorno nomeado\n\n")
	soma, subtracao := calcMath(10, 2)
	fmt.Println("Resultado", soma, subtracao)
}
