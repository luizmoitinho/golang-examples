package main

import "fmt"

var n int

func init() {
	fmt.Println("1º Funcao init")
	n = 10
}

func main() {
	fmt.Println("2º Funcao Main")
	fmt.Println("Valor de n = ", n)
}
