package enderecos

import "testing"

/*
	 	- Criar um arquivo de com seu nome finalizando em: "_test"
		- Criar com no mesmo pacote da função a ser testada
		- Nome da função deve iniciar coim a palavra "Test"+"NomeFuncao"
		- Parametro especifico = Ponteiro do tipo T
*/

func TestTipoEndereco(t *testing.T) {
	enderecoTeste := "Avenida Paulista"
	tipoEnderecoEsperado := "Avenida"

	if TipoEndereco(enderecoTeste) != tipoEnderecoEsperado {
		t.Error("O tipo recebido é diferente do esperado.")
	}
	
}
