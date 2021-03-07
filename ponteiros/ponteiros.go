package main

import "fmt"

func main() {
	fmt.Println("Ponteiros")
	var var1 int = 10
	var pVar1 *int

	fmt.Println("Valor de var1 = ", var1)
	fmt.Println("Endereco de var1 = ", &var1)

	fmt.Println("Endereco de pVar1 = ", pVar1)

	pVar1 = &var1
	fmt.Println("pVar1 = &var1: pVar1 = (", *pVar1, pVar1, ")")

	*pVar1 = 100
	fmt.Println("*pVar1 =", *pVar1)
	fmt.Println("var1 =", var1)
}
