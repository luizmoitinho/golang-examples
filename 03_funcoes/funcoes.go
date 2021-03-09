package main

import "fmt"

func main() {

	fmt.Println("10 + 2 = ", somar(10, 2))

	fmt.Printf("\nFuncao variavel\n")

	var msgBoasVindas = func(nome string) string {
		fmt.Println("Seja bem-vindo,", nome)
		return nome
	}

	mensagem := msgBoasVindas("Luiz Moitinho")
	fmt.Println(mensagem)

	fmt.Printf("\nFuncao retornando mais de um valor\n")
	soma, subtracao := calculos(19, 1)
	fmt.Println("Resultados: ", soma, subtracao)
}

func somar(n1 int8, n2 int8) int8 {

	return n1 + n2

}

func calculos(n1, n2 int8) (int8, int8) {

	return n1 + n2, n1 - n2

}
