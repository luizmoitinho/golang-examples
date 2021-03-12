package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Multiplexador")
	channel := multiplexar(escrever("Programando em Golang"), escrever("Padrões de concorrência"))
	for i := 0; i < 10; i++ {
		fmt.Println(<-channel)
	}
}

func multiplexar(chan1, chan2 <-chan string) <-chan string {
	outChannel := make(chan string)
	go func() {
		for {
			select {
			case mensagem := <-chan1:
				outChannel <- mensagem
			case mensagem := <-chan2:
				outChannel <- mensagem
			}
		}
	}()

	return outChannel
}

func escrever(text string) <-chan string {
	channel := make(chan string)
	go func() {
		for {
			channel <- fmt.Sprintf("valor recebido: %s", text)
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(2000)))
		}
	}()
	return channel
}
