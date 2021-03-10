package main

import "fmt"

func fibonacci(pos uint) uint {
	if pos <= 1 {
		return pos
	}
	return fibonacci(pos-2) + fibonacci(pos-1)

}

func main() {
	fmt.Println("Funções recursivas: fibonacci")
	posicao := uint(2)
	fmt.Println(fibonacci(posicao))
}
