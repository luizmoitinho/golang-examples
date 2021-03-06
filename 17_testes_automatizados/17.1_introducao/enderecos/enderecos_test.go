package enderecos_test

import (
	. "introducao-testes/enderecos"
	"testing"
)

type cenarioTeste struct {
	enderecoInserido string
	retornoEsperado  string
}

func TestTipoEndereco(t *testing.T) {

	t.Parallel()

	cenariosTeste := []cenarioTeste{
		{"Rua ABC", "Rua"},
		{"RUA ABC", "Rua"},
		{"Avenida Paulista", "Avenida"},
		{"AVENIDA REBOUÇAS", "Avenida"},
		{"Av. Paulista", "Av."},
		{"Av Paulista", "Av"},
		{"Rodovia dos Imigrantes", "Rodovia"},
		{"Estrada da Colônia", "Estrada"},
		{"Praça das Rosas", "Tipo inválido"},
		{"", "Tipo inválido"},
	}

	for _, cenario := range cenariosTeste {
		resultado := TipoEndereco(cenario.enderecoInserido)
		if resultado != cenario.retornoEsperado {
			t.Errorf("O tipo recebido %s é diferente do esperado %s", resultado, cenario.retornoEsperado)
		}
	}

}

func TestQualquer(t *testing.T) {
	t.Parallel()

	if 1 > 2 {
		t.Errorf("Teste falhou.")
	}
}
