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
	verdadeiro, falso := true, false

	fmt.Println("true && false: ", verdadeiro && falso)

	fmt.Println("true || false: ", verdadeiro || falso)

	fmt.Println("!true: ", !verdadeiro)

	fmt.Println("!false: ", !falso)

	fmt.Printf("\nOperadores unários: ++, -- \n")
	numero := 10
	numero++
	fmt.Println("numero = 10")
	fmt.Println("numero++:", numero)
	numero--
	fmt.Println("numero--:", numero)
	numero += 3
	fmt.Println("numero+=3:", numero)

}
