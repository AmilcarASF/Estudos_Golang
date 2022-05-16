package enderecos

import "testing"

type cenarioDeTeste struct {
	enderecoInserido string
	retornoEsperado  string
}

// Teste de Unidade

func TestTipoDeEndereco(t *testing.T) {

	cenariosDeTeste := []cenarioDeTeste{
		{"Rua da saudade", "Rua"},
		{"Avenida Paulista", "Avenida"},
		//{"Praça por do sol", "Tipo inválido"},
		{"Rua dos cristais", "Rua"},
		{"rua primeiro de maio", "Rua"},
		{"Rodovia dos Imigrantes", "Rodovia"},
		{"ESTRADA de areia", "Estrada"},
		//{"", "Tipo inválido"},
	}

	for _, cenario := range cenariosDeTeste {
		retornoRecebido := TipoDeEndereco(cenario.enderecoInserido)
		if retornoRecebido != cenario.retornoEsperado {
			t.Errorf("O tip recebido %s é diferente do esperado %s",
				retornoRecebido, cenario.retornoEsperado)
		}
	}
}
