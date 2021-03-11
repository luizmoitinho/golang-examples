package main

import "fmt"

func main() {
	channel := make(chan string, 2)

	channel <- "Canal 1"
	channel <- "Canal 2"
	msg1 := <-channel
	msg2 := <-channel

	fmt.Println(msg1, msg2)
}
