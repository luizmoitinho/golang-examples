package main

import "fmt"

type Usuario struct {
	nome  string
	idade uint8
}

func (u Usuario) salvar() {
	fmt.Println("Salvando usuario: ", u.nome)
}

func main() {
	fmt.Println("# Métodos")
	u1 := Usuario{"Luiz Moitinho", 22}
	u1.salvar()
}
