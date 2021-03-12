package enderecos

import (
	"strings"
)

//TipoEndereco: Verifica se um endereço tem o tipo valido e retorna
func TipoEndereco(endereco string) string {
	tiposValidos := []string{"Rua", "Avenida", "Av.", "Av", "Estrada", "Rodovia", "Viela"}

	primeiraPalavra := strings.Split(strings.ToLower(endereco), " ")[0]
	isEnderecoValido := false
	for _, tipo := range tiposValidos {
		if strings.ToLower(tipo) == primeiraPalavra {
			isEnderecoValido = true
			break
		}
	}

	if isEnderecoValido {
		return strings.Title(primeiraPalavra)
	}
	return "Tipo inválido"

}
