package main

import "fmt"

//Usuario ...
type Usuario struct {
	Nome     string
	Idade    uint8
	Endereco Endereco
}

//Endereco ...
type Endereco struct {
	logradouro string
	numero     uint8
}

func main() {
	fmt.Println("Arquivo structs")
	var u Usuario
	displayUser(u)
	u.Nome = "Luiz"
	u.Idade = 22
	displayUser(u)

	endereco := Endereco{"Rua das Flores", 20}
	u2 := Usuario{"Carlos", 66, endereco}
	displayUser(u2)

	u3 := Usuario{
		Nome: "João",
	}
	displayUser(u3)

}

func displayUser(u Usuario) {
	fmt.Printf("\nDados de Usuario\n")
	fmt.Println("Nome: ", u.Nome)
	fmt.Println("Idade:", u.Idade)
	fmt.Println("Endereco:", u.Endereco)
}
