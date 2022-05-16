package main

import (
	"fmt"
	"testes_automatiado/enderecos"
)

func main() {
	fmt.Println("Testes Automaziados")

	tipoEndereco := enderecos.TipoDeEndereco("Rodovia da Saudade")
	fmt.Println(tipoEndereco)
}
