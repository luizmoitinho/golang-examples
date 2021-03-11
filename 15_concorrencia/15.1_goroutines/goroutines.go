package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Go Routines")
	go escrever("Olá mundo!")
	escrever("Programando em Go!")
}

func escrever(texto string) {
	for {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
