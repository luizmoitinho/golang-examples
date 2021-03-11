package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Channel")

	channel := make(chan string)
	go escrever("Go routine 1", channel)

	for {
		msg, open := <-channel
		if !open {
			break
		}
		fmt.Println(msg)
	}

}

func escrever(texto string, channel chan string) {
	for i := 0; i < 5; i++ {
		channel <- texto
		time.Sleep(time.Second)
	}
	close(channel)
}
