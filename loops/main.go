package main

import (
	"fmt"
)

func main() {
	fmt.Println("Loops")

	i := 0

	for i < 10 {
		i++
		fmt.Println("i", i)
		//time.Sleep(time.Second)
	}

	for j := 0; j < 10; j++ {
		fmt.Println("j", j)
	}

	nomes := []string{"Amilcar", "Camila", "Gabriel"}

	for indice, nome := range nomes {
		fmt.Println(indice, nome)
	}

	usuario := map[string]string{
		"nome":      "Flora",
		"sobreNome": "Ferreira",
	}

	fmt.Println(usuario)

	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}
}
