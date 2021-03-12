package main

import "fmt"

func main() {
	fmt.Println("Worker Pools")
	tarefas := make(chan uint, 45)
	resultado := make(chan uint, 45)

	go worker(tarefas, resultado)
	go worker(tarefas, resultado)
	go worker(tarefas, resultado)
	go worker(tarefas, resultado)
	go worker(tarefas, resultado)
	go worker(tarefas, resultado)
	go worker(tarefas, resultado)
	go worker(tarefas, resultado)

	for i := uint(0); i < 45; i++ {
		tarefas <- i
	}
	close(tarefas)

	for i := uint(0); i < 45; i++ {
		resultado := <-resultado
		fmt.Println(resultado)
	}

}

func worker(tarefas <-chan uint, resultado chan<- uint) {

	for numero := range tarefas {
		resultado <- fibonacci(numero)
	}

}

func fibonacci(pos uint) uint {
	if pos <= 1 {
		return pos
	}
	return fibonacci(pos-2) + fibonacci(pos-1)

}
