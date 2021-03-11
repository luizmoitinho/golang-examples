package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("Wait Group")

	var waitGroup sync.WaitGroup

	waitGroup.Add(2)

	go func() {
		escrever("Go routine 1")
		waitGroup.Done() // -1
	}()

	go func() {
		escrever("Go routine 2")
		waitGroup.Done() // -1
	}()

	waitGroup.Wait() // 0

	fmt.Println("# Fim do programa")
}

func escrever(texto string) {
	for i := 0; i < 5; i++ {
		fmt.Println(i, "º:", texto)
		time.Sleep(time.Second)
	}
}
