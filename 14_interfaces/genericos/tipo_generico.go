package main

import "fmt"

func generica(i interface{}) {
	fmt.Println(i)
}

func main() {
	fmt.Println("Tipo Genéricos")
	generica("string")
	generica(100)
	generica(true)

}
