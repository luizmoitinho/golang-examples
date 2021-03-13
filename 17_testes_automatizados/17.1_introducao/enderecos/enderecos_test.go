package enderecos

import "testing"

/*
	 	- Criar um arquivo de com seu nome finalizando em: "_test"
		- Criar com no mesmo pacote da função a ser testada
		- Nome da função deve iniciar coim a palavra "Test"+"NomeFuncao"
		- Parametro especifico = Ponteiro do tipo T
*/

func TestTipoEndereco(t *testing.T) {
	enderecoTeste := "Rua Paulista"
	tipoEnderecoEsperado := "Avenida"
	resultado := TipoEndereco(enderecoTeste)

	if resultado != tipoEnderecoEsperado {
		t.Errorf("O tipo recebido é diferente do esperado. Esperava %s e recebeu %s",
			tipoEnderecoEsperado, resultado)
	}

}
