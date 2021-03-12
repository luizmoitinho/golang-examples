package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Generator")
	channel := escrever("Programando em Golang")
	for i := 0; i < 10; i++ {
		fmt.Println(<-channel)
	}
}

func escrever(text string) <-chan string {
	channel := make(chan string)
	go func() {
		for {
			channel <- fmt.Sprintf("valor recebido: %s", text)
			time.Sleep(time.Millisecond * 500)
		}
	}()
	return channel
}
