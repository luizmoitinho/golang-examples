package main

import "fmt"

type Usuario struct {
	nome   string
	idade  uint8
	status bool
}

func (u *Usuario) salvar() {
	u.status = true
	fmt.Println("Salvando usuario: ", u.nome)
}

func (u *Usuario) isAdulto() bool {
	return u.idade >= 18
}
func main() {
	fmt.Println("# Métodos")

	u1 := Usuario{"Luiz Moitinho", 22, false}

	u1.salvar()
	fmt.Println(u1)

	fmt.Println("Usuario ", u1.nome, "é adulto?", u1.isAdulto())

}
