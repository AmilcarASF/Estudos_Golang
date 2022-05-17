package main

import (
	"bytes"
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
	usuario1 := usuario{"Amilcar", "Rua dos cristais", 35}
	fmt.Println(usuario1)

	usuarioJson, erro := json.Marshal(usuario1)
	if erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(bytes.NewBuffer(usuarioJson))
}
