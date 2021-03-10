package main

import "fmt"

func inverterSinal(num *int64) {
	*num = *num * -1
}

func main() {
	numero := int64(20)
	inverterSinal(&numero)
	fmt.Println(numero)

}
