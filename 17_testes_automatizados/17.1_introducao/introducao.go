package main

import (
	"fmt"
	"introducao-testes/enderecos"
)

func main() {
	tipoEndereco := enderecos.TipoEndereco("Avenida Paulista")
	fmt.Println(tipoEndereco)
	tipoEndereco = enderecos.TipoEndereco("Rodovia dos imigrantes")
	fmt.Println(tipoEndereco)
}
