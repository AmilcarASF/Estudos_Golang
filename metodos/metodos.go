package main

import "fmt"

type usuario struct {
	nome  string
	idade uint8
}

func (u usuario) salvar() {
	fmt.Printf("Salvando os dados do Usuário %s no banco de dados\n", u.nome)
}

func (u *usuario) fazerAniversario() {
	u.idade++
}

func main() {
	fmt.Println("Métodos")

	usuario1 := usuario{"Amilcar", 35}
	fmt.Println(usuario1)
	usuario1.salvar()
	usuario1.fazerAniversario()
	fmt.Println(usuario1.idade)

}
