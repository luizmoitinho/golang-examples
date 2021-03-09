package main

import "fmt"

//Pessoa ...
type Pessoa struct {
	nome      string
	sobrenome string
	idade     uint8
	altura    uint8
}

//Estudante ...
type Estudante struct {
	Pessoa
	matricula uint
	curso     string
	faculdade string
}

func main() {
	fmt.Println("Herança")
	p1 := Pessoa{
		"Luiz", "Carlos", 20, 170,
	}
	fmt.Println("Pessoa: ", p1)

	e1 := Estudante{p1, 123, "Sistemas de Informação", "UFS"}
	fmt.Println("Estudante: ", e1)

	fmt.Println("Nome Estudante: ", e1.Pessoa.nome)
	fmt.Println("Sobrenome Estudante: ", e1.sobrenome)
}
