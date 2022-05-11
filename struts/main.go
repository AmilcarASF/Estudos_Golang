package main

import "fmt"

type usuario struct {
	nome  string
	idade uint8
	endereco
}

type endereco struct {
	logradouro string
	numero     int8
}

func main() {

	enderecoExemplo := endereco{"Rua da saudade", 123}

	var u usuario
	u.nome = "Amilcar"
	u.idade = 34
	fmt.Println(u)

	u2 := usuario{nome: "Gabriel", idade: 6}
	fmt.Println(u2)

	u3 := usuario{nome: "Flora"}
	fmt.Println(u3)

	u4 := usuario{"José", 80, enderecoExemplo}
	fmt.Println(u4)

}
