package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type usuario struct {
	Nome     string `json:"nome"`
	Endereco string `json:"logradouro"`
	Idade    uint   `json:"idade"`
}

func main() {
	usuarioEmJSON := `{"nome":"Amilcar","logradouro":"Rua dos cristais","idade":35}`

	var usuario1 usuario

	if erro := json.Unmarshal([]byte(usuarioEmJSON), &usuario1); erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(usuario1)
}
