package main

import "fmt"

func main() {
	//int8, int16, int32 e int64
	fmt.Println("Tipos de inteiros: int8, int16, int32 e int64 e ")
	var num int8 = 100
	fmt.Println("Int16: ", num)

	fmt.Println("Int sem sinal: unsigned")
	var unsingned uint16 = 100
	fmt.Println("uint16: ", unsingned)

	//alias
	//int32 = rune
	var ASCII rune = 250
	fmt.Printf("Rune: %c\n", ASCII)

	//int8 = byte
	var bits byte = 8
	fmt.Printf("bits: %d\n", bits)

	fmt.Println("============================")
	fmt.Println("Tipos de ponto flutuante: ")
	var numReal1 float32 = 123.5
	fmt.Println("float32: ", numReal1)

	var numReal2 float64 = 12300000000000.5
	fmt.Println("float64: ", numReal2)

}
