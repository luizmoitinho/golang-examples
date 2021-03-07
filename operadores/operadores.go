package main

import "fmt"

func main() {
	//Aritméticos: +, -, *, / e %
	fmt.Printf("\nOperadores aritméticos\n")
	soma := 10 + 2
	subtracao := 10 - 2
	divisao := 10 / 2
	multiplicacao := 10 * 2
	restoDivisao := 10 % 2

	fmt.Println(soma, subtracao, divisao, multiplicacao, restoDivisao)

	fmt.Printf("\nOperadores relacionais: <, >, >=, <=, ==, !=\n")
	fmt.Println(1 > 2)
	fmt.Println(1 == 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 < 2)
	fmt.Println(1 != 2)

	fmt.Printf("\nOperadores relacionais: &&, ||, ! \n")
}
