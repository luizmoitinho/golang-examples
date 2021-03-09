package main

import "fmt"

func main() {

	msg := func(msg string) string {
		return fmt.Sprintf("Mensagem: %s", msg)
	}("Dentro de uma função anonima")

	fmt.Println(msg)
}
